package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/ghodss/yaml"
)

// metadata is a combination of read and derived metadata.
type metadata struct {
	Resources []*entry `json:"resources"`
}

// entry in a metadata file
type entry struct {
	Kind           string   `json:"kind"`
	ListKind       string   `json:"listKind"`
	Singular       string   `json:"singular"`
	Plural         string   `json:"plural"`
	Group          string   `json:"group"`
	Versions       []string `json:"versions"`
	Proto          string   `json:"proto"`
	Converter      string   `json:"converter"`
	ProtoGoPackage string   `json:"protoPackage"`
	Collection     string   `json:"collection"`
	Generated      string   `json:"generated"`
	Optional       bool     `json:"optional"`
}

func main() {
	input := os.Args[1]
	output := os.Args[2]

	m, err := readMetadata(input)
	if err != nil {
		fmt.Printf("Error reading metadata: %v", err)
		os.Exit(-2)
	}

	groupTpl, err := template.ParseGlob("templates/group/*.go")
	if err != nil {
		fmt.Printf("parsing group: %v", err)
		os.Exit(-3)
	}
	groupTpl.Funcs(sprig.FuncMap())

	versionTpl, err := template.ParseGlob("templates/version/*.go")
	if err != nil {
		fmt.Printf("parsing version: %v", err)
		os.Exit(-3)
	}
	versionTpl.Funcs(sprig.FuncMap())

	typeTpl, err := template.ParseFiles("templates/type.go")
	if err != nil {
		fmt.Printf("parsing type: %v", err)
		os.Exit(-3)
	}
	typeTpl.Funcs(sprig.FuncMap())

	groups := map[string][]int{}

	for i, resource := range m.Resources {
		if resource.Group == "" {
			continue
		}
		name := strings.TrimSuffix(resource.Group, ".istio.io")
		groups[name] = append(groups[name], i)
	}

	for group, ids := range groups {
		ctx := struct {
			Group        string
			Version      string
			Type         string
			IstioPackage string
			Message      string
		}{
			Group: group,
		}

		if err := ensureDir(filepath.Join(output, group)); err != nil {
			fmt.Println(err)
		}

		for _, tpl := range groupTpl.Templates() {
			if err := ensureTemplate(ctx, tpl, filepath.Join(output, group, tpl.Name())); err != nil {
				fmt.Println(err)
			}
		}

		for _, id := range ids {
			resource := m.Resources[id]
			if strings.Contains(resource.Proto, "googleapis") {
				continue
			}
			ctx.Type = strings.Title(resource.Kind)
			parts := strings.Split(resource.Proto, ".")
			ctx.Message = parts[len(parts)-1]
			ctx.IstioPackage = strings.Join(parts[1:len(parts)-1], "/")

			for _, version := range resource.Versions {
				ctx.Version = version

				if err := ensureDir(filepath.Join(output, group, version)); err != nil {
					fmt.Println(err)
				}

				// can optimize this
				for _, tpl := range versionTpl.Templates() {
					if err := ensureTemplate(ctx, tpl, filepath.Join(output, group, version, tpl.Name())); err != nil {
						fmt.Println(err)
					}
				}

				if err := ensureTemplate(ctx, typeTpl, filepath.Join(output, group, version, strings.ToLower(resource.Kind)+"_type.go")); err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	if err != nil {
		fmt.Printf("Error applying template: %v", err)
		os.Exit(-3)
	}
}

func readMetadata(path string) (*metadata, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file: %v", err)
	}

	var m metadata

	if err = yaml.Unmarshal(b, &m); err != nil {
		return nil, fmt.Errorf("error marshalling input file: %v", err)
	}

	// Auto-complete listkind fields with defaults.
	for _, item := range m.Resources {
		if item.ListKind == "" {
			item.ListKind = item.Kind + "List"
		}
	}

	return &m, nil
}

func ensureTemplate(ctx interface{}, tpl *template.Template, target string) error {
	var b bytes.Buffer
	if err := tpl.Execute(&b, ctx); err != nil {
		return err
	}

	if err := ioutil.WriteFile(target, b.Bytes(), 0664); err != nil {
		return err
	}

	return nil
}

func getID(collection string) string {
	out := ""

	// Convert to camelcase by capitalizing the first letter in each word separated by "/".
	for _, part := range strings.Split(collection, "/") {
		out += strings.Title(part)
	}

	return out
}

func ensureDir(path string) error {
	i, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(path, os.ModePerm)
		}

		return err
	}

	if !i.IsDir() {
		return fmt.Errorf("not a dir")
	}

	return nil
}
