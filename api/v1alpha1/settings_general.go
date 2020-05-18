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
	parameters "k8s.libre.sh/application/settings/parameters"
)

type General struct {
	GeneralConfig  `json:",inline"`
	GeneralSecrets `json:",inline"`
}

type GeneralConfig struct {
	URL        parameters.Parameter `json:"url,omitempty" env:"ROOT_URL"`
	InstanceIP parameters.Parameter `json:"instanceIP,omitempty" env:"INSTANCE_IP"`
	Username   parameters.Parameter `json:"password,omitempty" env:"ADMIN_USERNAME"`
}

type GeneralSecrets struct {
	Password parameters.Parameter `json:"username,omitempty" env:"ADMIN_PASS"`
}

func (s *General) SetDefaults() {
	if len(s.InstanceIP.FromKey) == 0 {
		s.InstanceIP.FromKey = "status.podIP"
	}
	if len(s.InstanceIP.Ref) == 0 {
		s.InstanceIP.Ref = "v1"
	}
	if len(s.InstanceIP.Type) == 0 {
		s.InstanceIP.Type = "field"
	}
}

func (s *General) GetParameters() parameters.Parameters {
	// TODO manage error
	params, _ := parameters.Marshal(s.GeneralConfig)
	secretParams, _ := parameters.Marshal(s.GeneralSecrets)

	for _, p := range secretParams {
		if len(p.Type) == 0 {
			p.Type = parameters.SecretParameter
		}
		if len(p.ValueFrom.Ref) == 0 && len(p.Generate) == 0 && len(p.Value) == 0 {
			p.Generate = parameters.GenerateRand12
		}
	}

	params = append(params, secretParams...)

	return params
}
