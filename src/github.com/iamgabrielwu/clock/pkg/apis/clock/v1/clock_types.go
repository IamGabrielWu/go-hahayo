package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ClockSpec defines the desired state of Clock
type ClockSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Reset       bool   `json:"reset"`
	CurrentDay  string `json:"current_day"`
	CurrentTime string `json:"current_time"`
}

// ClockStatus defines the observed state of Clock
type ClockStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Reset       bool   `json:"reset"`
	CurrentDay  string `json:"current_day"`
	CurrentTime string `json:"current_time"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Clock is the Schema for the clocks API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=clocks,scope=Namespaced
type Clock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClockSpec   `json:"spec,omitempty"`
	Status ClockStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClockList contains a list of Clock
type ClockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Clock `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Clock{}, &ClockList{})
}
