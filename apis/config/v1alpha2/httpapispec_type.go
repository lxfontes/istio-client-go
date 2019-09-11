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

// HTTPAPISpec defines the desired state of HTTPAPISpec
type HTTPAPISpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec HTTPAPISpecSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPAPISpecList is a list of HTTPAPISpec resources
type HTTPAPISpecList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []HTTPAPISpec `json:"items"`
}

// GetSpecMessage ...
func (o *HTTPAPISpec) GetSpecMessage() proto.Message {
	return &o.Spec.HTTPAPISpec
}

// +k8s:deepcopy-gen=false

// HTTPAPISpecSpec ...
type HTTPAPISpecSpec struct {
	istio.HTTPAPISpec
}

// MarshalJSON ...
func (p *HTTPAPISpecSpec) MarshalJSON() ([]byte, error) {
	buffer := bytes.Buffer{}
	writer := bufio.NewWriter(&buffer)
	marshaler := jsonpb.Marshaler{}
	if err := marshaler.Marshal(writer, &p.HTTPAPISpec); err != nil {
		return nil, err
	}

	writer.Flush()
	return buffer.Bytes(), nil
}

// UnsmarshalJSON ...
func (p *HTTPAPISpecSpec) UnmarshalJSON(b []byte) error {
	reader := bytes.NewReader(b)
	unmarshaler := jsonpb.Unmarshaler{}
	if err := unmarshaler.Unmarshal(reader, &p.HTTPAPISpec); err != nil {
		return err
	}
	return nil
}


// DeepCopyInto is a deepcopy function, copying the receiver, writing into out. in must be non-nil.
// Based of https://github.com/istio/istio/blob/release-0.8/pilot/pkg/config/kube/crd/types.go#L450
func (in *HTTPAPISpecSpec) DeepCopyInto(out *HTTPAPISpecSpec) {
	*out = *in
}

func init() {
	scheme.AddKnownTypes(SchemeGroupVersion, &HTTPAPISpec{}, &HTTPAPISpecList{})
}
