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
	components "k8s.libre.sh/components"
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
}

// +kubebuilder:object:root=true

// +kubebuilder:resource:shortName=rc
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="version",type="string",JSONPath=".spec.Version",description="rocketchat version"
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
