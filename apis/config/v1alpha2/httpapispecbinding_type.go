package v1alpha2

import (
	"bufio"
	"bytes"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	istio "istio.io/api/mixer/v1/config/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// // +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPAPISpecBinding defines the desired state of HTTPAPISpecBinding
type HTTPAPISpecBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec HTTPAPISpecBindingSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPAPISpecBindingList is a list of HTTPAPISpecBinding resources
type HTTPAPISpecBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []HTTPAPISpecBinding `json:"items"`
}

// GetSpecMessage ...
func (o *HTTPAPISpecBinding) GetSpecMessage() proto.Message {
	return &o.Spec.HTTPAPISpecBinding
}

// +k8s:deepcopy-gen=false

// HTTPAPISpecBindingSpec ...
type HTTPAPISpecBindingSpec struct {
	istio.HTTPAPISpecBinding
}

// MarshalJSON ...
func (p *HTTPAPISpecBindingSpec) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	writer := bufio.NewWriter(&buffer)
	marshaler := jsonpb.Marshaler{}
	if err := marshaler.Marshal(writer, &p.HTTPAPISpecBinding); err != nil {
		return nil, err
	}

	writer.Flush()
	return buffer.Bytes(), nil
}

// UnsmarshalJSON ...
func (p *HTTPAPISpecBindingSpec) UnmarshalJSON(b []byte) error {
	reader := bytes.NewReader(b)
	unmarshaler := jsonpb.Unmarshaler{}
	if err := unmarshaler.Unmarshal(reader, &p.HTTPAPISpecBinding); err != nil {
		return err
	}
	return nil
}


// DeepCopyInto is a deepcopy function, copying the receiver, writing into out. in must be non-nil.
// Based of https://github.com/istio/istio/blob/release-0.8/pilot/pkg/config/kube/crd/types.go#L450
func (in *HTTPAPISpecBindingSpec) DeepCopyInto(out *HTTPAPISpecBindingSpec) {
	*out = *in
}

func init() {
	scheme.AddKnownTypes(SchemeGroupVersion, &HTTPAPISpecBinding{}, &HTTPAPISpecBindingList{})
}
