package v1alpha3

import (
	"bufio"
	"bytes"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	istio "istio.io/api/networking/v1alpha3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// // +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EnvoyFilter defines the desired state of EnvoyFilter
type EnvoyFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec EnvoyFilterSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EnvoyFilterList is a list of EnvoyFilter resources
type EnvoyFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EnvoyFilter `json:"items"`
}

// GetSpecMessage ...
func (o *EnvoyFilter) GetSpecMessage() proto.Message {
	return &o.Spec.EnvoyFilter
}

// +k8s:deepcopy-gen=false

// EnvoyFilterSpec ...
type EnvoyFilterSpec struct {
	istio.EnvoyFilter
}

// MarshalJSON ...
func (p *EnvoyFilterSpec) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	writer := bufio.NewWriter(&buffer)
	marshaler := jsonpb.Marshaler{}
	if err := marshaler.Marshal(writer, &p.EnvoyFilter); err != nil {
		return nil, err
	}

	writer.Flush()
	return buffer.Bytes(), nil
}

// UnsmarshalJSON ...
func (p *EnvoyFilterSpec) UnmarshalJSON(b []byte) error {
	reader := bytes.NewReader(b)
	unmarshaler := jsonpb.Unmarshaler{}
	if err := unmarshaler.Unmarshal(reader, &p.EnvoyFilter); err != nil {
		return err
	}
	return nil
}


// DeepCopyInto is a deepcopy function, copying the receiver, writing into out. in must be non-nil.
// Based of https://github.com/istio/istio/blob/release-0.8/pilot/pkg/config/kube/crd/types.go#L450
func (in *EnvoyFilterSpec) DeepCopyInto(out *EnvoyFilterSpec) {
	*out = *in
}

func init() {
	scheme.AddKnownTypes(SchemeGroupVersion, &EnvoyFilter{}, &EnvoyFilterList{})
}
