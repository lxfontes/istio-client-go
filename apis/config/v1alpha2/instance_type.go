package v1alpha2

import (
	"bufio"
	"bytes"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	istio "istio.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// // +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Instance defines the desired state of Instance
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec InstanceSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InstanceList is a list of Instance resources
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Instance `json:"items"`
}

// GetSpecMessage ...
func (o *Instance) GetSpecMessage() proto.Message {
	return &o.Spec.Instance
}

// +k8s:deepcopy-gen=false

// InstanceSpec ...
type InstanceSpec struct {
	istio.Instance
}

// MarshalJSON ...
func (p *InstanceSpec) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	writer := bufio.NewWriter(&buffer)
	marshaler := jsonpb.Marshaler{}
	if err := marshaler.Marshal(writer, &p.Instance); err != nil {
		return nil, err
	}

	writer.Flush()
	return buffer.Bytes(), nil
}

// UnsmarshalJSON ...
func (p *InstanceSpec) UnmarshalJSON(b []byte) error {
	reader := bytes.NewReader(b)
	unmarshaler := jsonpb.Unmarshaler{}
	if err := unmarshaler.Unmarshal(reader, &p.Instance); err != nil {
		return err
	}
	return nil
}


// DeepCopyInto is a deepcopy function, copying the receiver, writing into out. in must be non-nil.
// Based of https://github.com/istio/istio/blob/release-0.8/pilot/pkg/config/kube/crd/types.go#L450
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
}

func init() {
	scheme.AddKnownTypes(SchemeGroupVersion, &Instance{}, &InstanceList{})
}
