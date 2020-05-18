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
	settings "k8s.libre.sh/application/settings"
	meta "k8s.libre.sh/meta"
	objects "k8s.libre.sh/objects"
	"k8s.libre.sh/utils"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Settings struct {
	*meta.ObjectMeta `json:"meta,omitempty"`
	SecretMeta       *meta.ObjectMeta `json:"secretMeta,omitempty"`
	ConfigMeta       *meta.ObjectMeta `json:"configMeta,omitempty"`
	ConfigRefs       []string         `json:"configRefs,omitempty"`
	SecretRefs       []string         `json:"secretRefs,omitempty"`
	Database         Database         `json:"database,omitempty"`
	SMTP             SMTP             `json:"smtp,omitempty"`
	General          General          `json:"general,omitempty"`
}

func (s *Settings) GetMeta() meta.Instance { return s.ObjectMeta }

func (s *Settings) SetDefaults() {
	if s.ObjectMeta == nil {
		s.ObjectMeta = new(meta.ObjectMeta)
	}

	if len(s.ObjectMeta.Labels) == 0 {
		s.ObjectMeta.Labels = make(map[string]string)
	}

	s.ObjectMeta.Labels["app.kubernetes.io/component"] = "settings"

	//	meta.SetObjectMeta(i, s.ObjectMeta)

	s.General.SetDefaults()
	s.Database.SetDefaults()
	s.SMTP.SetDefaults()
}

func (s *Settings) GetConfig() settings.Config {

	params := s.General.GetParameters()
	params = append(params, s.Database.GetParameters()...)
	params = append(params, s.SMTP.GetParameters()...)

	secretRefs := s.SecretRefs
	configRefs := s.ConfigRefs

	// TO FIX OBJECTMETA
	s.ObjectMeta.Labels["app.kubernetes.io/component"] = "settings"

	settings := &settings.ConfigSpec{
		SecretRefs: utils.Unique(secretRefs),
		ConfigRefs: utils.Unique(configRefs),
		Parameters: &params,
	}

	return settings
}

func (s *Settings) GetSecretMeta() meta.Instance {
	if s.SecretMeta == nil {
		s.SecretMeta = new(meta.ObjectMeta)
	}
	return s.SecretMeta
}
func (s *Settings) GetConfigMapMeta() meta.Instance {
	if s.ConfigMeta == nil {
		s.ConfigMeta = new(meta.ObjectMeta)
	}
	return s.ConfigMeta
}

func (s *Settings) GetObjects() map[int]objects.Object {

	return nil
}

func (s *Settings) Init(c client.Client) error {

	return nil
}
