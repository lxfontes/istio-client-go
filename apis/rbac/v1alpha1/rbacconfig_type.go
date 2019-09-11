package v1alpha1

import (
	"bufio"
	"bytes"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	istio "istio.io/api/rbac/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// // +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RbacConfig defines the desired state of RbacConfig
type RbacConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RbacConfigSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RbacConfigList is a list of RbacConfig resources
type RbacConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RbacConfig `json:"items"`
}

// GetSpecMessage ...
func (o *RbacConfig) GetSpecMessage() proto.Message {
	return &o.Spec.RbacConfig
}

// +k8s:deepcopy-gen=false

// RbacConfigSpec ...
type RbacConfigSpec struct {
	istio.RbacConfig
}

// MarshalJSON ...
func (p *RbacConfigSpec) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	writer := bufio.NewWriter(&buffer)
	marshaler := jsonpb.Marshaler{}
	if err := marshaler.Marshal(writer, &p.RbacConfig); err != nil {
		return nil, err
	}

	writer.Flush()
	return buffer.Bytes(), nil
}

// UnsmarshalJSON ...
func (p *RbacConfigSpec) UnmarshalJSON(b []byte) error {
	reader := bytes.NewReader(b)
	unmarshaler := jsonpb.Unmarshaler{}
	if err := unmarshaler.Unmarshal(reader, &p.RbacConfig); err != nil {
		return err
	}
	return nil
}


// DeepCopyInto is a deepcopy function, copying the receiver, writing into out. in must be non-nil.
// Based of https://github.com/istio/istio/blob/release-0.8/pilot/pkg/config/kube/crd/types.go#L450
func (in *RbacConfigSpec) DeepCopyInto(out *RbacConfigSpec) {
	*out = *in
}

func init() {
	scheme.AddKnownTypes(SchemeGroupVersion, &RbacConfig{}, &RbacConfigList{})
}
