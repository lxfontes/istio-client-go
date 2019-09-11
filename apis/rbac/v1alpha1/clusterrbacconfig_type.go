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

// ClusterRbacConfig defines the desired state of ClusterRbacConfig
type ClusterRbacConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterRbacConfigSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterRbacConfigList is a list of ClusterRbacConfig resources
type ClusterRbacConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ClusterRbacConfig `json:"items"`
}

// GetSpecMessage ...
func (o *ClusterRbacConfig) GetSpecMessage() proto.Message {
	return &o.Spec.RbacConfig
}

// +k8s:deepcopy-gen=false

// ClusterRbacConfigSpec ...
type ClusterRbacConfigSpec struct {
	istio.RbacConfig
}

// MarshalJSON ...
func (p *ClusterRbacConfigSpec) MarshalJSON() ([]byte, error) {
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
func (p *ClusterRbacConfigSpec) UnmarshalJSON(b []byte) error {
	reader := bytes.NewReader(b)
	unmarshaler := jsonpb.Unmarshaler{}
	if err := unmarshaler.Unmarshal(reader, &p.RbacConfig); err != nil {
		return err
	}
	return nil
}


// DeepCopyInto is a deepcopy function, copying the receiver, writing into out. in must be non-nil.
// Based of https://github.com/istio/istio/blob/release-0.8/pilot/pkg/config/kube/crd/types.go#L450
func (in *ClusterRbacConfigSpec) DeepCopyInto(out *ClusterRbacConfigSpec) {
	*out = *in
}

func init() {
	scheme.AddKnownTypes(SchemeGroupVersion, &ClusterRbacConfig{}, &ClusterRbacConfigList{})
}