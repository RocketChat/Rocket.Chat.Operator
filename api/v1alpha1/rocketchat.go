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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	components "k8s.libre.sh/application/components"
)

// Constants for condition
const (
	// Ready => controller considers this resource Ready
	Ready = "Ready"
	// Qualified => functionally tested
	Qualified = "Qualified"
	// Settled => observed generation == generation + settled means controller is done acting functionally tested
	Settled = "Settled"
	// Cleanup => it is set to track finalizer failures
	Cleanup = "Cleanup"
	// Error => last recorded error
	Error = "Error"

	ReasonInit = "Init"
)

// RocketchatSpec defines the desired state of Rocketchat
type RocketchatSpec struct {
	Version  string    `json:"version,omitempty"`
	Settings *Settings `json:"settings,omitempty"`
	Storage  Storage   `json:"storage,omitempty"`
	App      *App      `json:"app,omitempty"`
}

type App struct {
	*components.Workload `json:",inline"`
}

// RocketchatStatus defines the observed state of Rocketchat
type RocketchatStatus struct {
	// ObservedGeneration is the most recent generation observed. It corresponds to the
	// Object's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// Conditions represents the latest state of the object
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,10,rep,name=conditions"`
	// Resources embeds a list of object statuses
	// +optional
	ComponentList `json:",inline,omitempty"`
	// ComponentsReady: status of the components in the format ready/total
	// +optional
	ComponentsReady string `json:"componentsReady,omitempty"`
}

// ObjectStatus is a generic status holder for objects
type ObjectStatus struct {
	// Link to object GetSelfLink()
	Link string `json:"link,omitempty"`
	// Name of object
	Name string `json:"name,omitempty"`
	// Kind of object
	Kind string `json:"kind,omitempty"`
	// Object group
	Group string `json:"group,omitempty"`
	// Status. Values: InProgress, Ready, Unknown
	Status string `json:"status,omitempty"`
}

// ConditionType encodes information on the condition
type ConditionType string

// Condition describes the state of an object at a certain point.
type Condition struct {
	// Type of condition.
	Type ConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=StatefulSetConditionType"`
	// Status of the condition, one of True, False, Unknown.
	//	Status corev1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus"`
	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,4,opt,name=reason"`
	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,5,opt,name=message"`
	// Last time the condition was probed
	// +optional
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty" protobuf:"bytes,3,opt,name=lastProbeTime"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,3,opt,name=lastTransitionTime"`
}

// ComponentList is a generic status holder for the top level resource
type ComponentList struct {
	// Object status array for all matching objects
	Objects []ObjectStatus `json:"components,omitempty"`
}

// +kubebuilder:object:root=true

// +kubebuilder:resource:shortName=rc
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="version",type="string",JSONPath=".spec.Version",description="rocketchat version"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// Rocketchat is the Schema for the rocketchats API
type Rocketchat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RocketchatSpec   `json:"spec,omitempty"`
	Status RocketchatStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RocketchatList contains a list of Rocketchat
type RocketchatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rocketchat `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Rocketchat{}, &RocketchatList{})
}
