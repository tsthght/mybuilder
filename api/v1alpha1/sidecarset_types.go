/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SidecarSetSpec defines the desired state of SidecarSet
type SidecarSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of SidecarSet. Edit SidecarSet_types.go to remove/update
	//Foo string `json:"foo,omitempty"`
	Selector   *metav1.LabelSelector `json:"selector,omitempty"`
	Containers []SidecarContainer    `json:"containers,omitempty"`
}

// SidecarContainer defines the container of Sidecar
type SidecarContainer struct {
	corev1.Container `json:""`
}

// SidecarSetStatus defines the observed state of SidecarSet
type SidecarSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	MatchedPods int32 `json:"matchedPods"`
	UpdatePods  int32 `json:"updatePods"`
	ReadyPods   int32 `json:"readyPods"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster

// SidecarSet is the Schema for the sidecarsets API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:Namespaced

// SidecarSet is the Schema for the sidecarsets API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="MATCHED",type="integer",JSONPath=".status.matchedPods",description="The number of pods matched."
// +kubebuilder:printcolumn:name="UPDATED",type="integer",JSONPath=".status.updatedPods",description="The number of pods matched and updated."
// +kubebuilder:printcolumn:name="READY",type="integer",JSONPath=".status.readyPods",description="The number of pods matched and ready."
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC."

type SidecarSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SidecarSetSpec   `json:"spec,omitempty"`
	Status SidecarSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SidecarSetList contains a list of SidecarSet
type SidecarSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SidecarSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SidecarSet{}, &SidecarSetList{})
}
