package {{ .Version }}

import (
	"bufio"
	"bytes"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	istio "istio.io/api/{{ .IstioPackage }}"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// // +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// {{ .Type }} defines the desired state of {{ .Type }}
type {{ .Type }} struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec {{ .Type }}Spec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// {{ .Type }}List is a list of {{ .Type }} resources
type {{ .Type }}List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []{{ .Type }} `json:"items"`
}

// GetSpecMessage ...
func (o *{{ .Type }}) GetSpecMessage() proto.Message {
	return &o.Spec.{{ .Message }}
}

// +k8s:deepcopy-gen=false

// {{ .Type }}Spec ...
type {{ .Type }}Spec struct {
	istio.{{ .Message }}
}

// MarshalJSON ...
func (p *{{ .Type }}Spec) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	writer := bufio.NewWriter(&buffer)
	marshaler := jsonpb.Marshaler{}
	if err := marshaler.Marshal(writer, &p.{{ .Message }}); err != nil {
		return nil, err
	}

	writer.Flush()
	return buffer.Bytes(), nil
}

// UnsmarshalJSON ...
func (p *{{ .Type }}Spec) UnmarshalJSON(b []byte) error {
	reader := bytes.NewReader(b)
	unmarshaler := jsonpb.Unmarshaler{}
	if err := unmarshaler.Unmarshal(reader, &p.{{ .Message }}); err != nil {
		return err
	}
	return nil
}


// DeepCopyInto is a deepcopy function, copying the receiver, writing into out. in must be non-nil.
// Based of https://github.com/istio/istio/blob/release-0.8/pilot/pkg/config/kube/crd/types.go#L450
func (in *{{ .Type }}Spec) DeepCopyInto(out *{{ .Type }}Spec) {
	*out = *in
}

func init() {
	scheme.AddKnownTypes(SchemeGroupVersion, &{{ .Type }}{}, &{{ .Type }}List{})
}
