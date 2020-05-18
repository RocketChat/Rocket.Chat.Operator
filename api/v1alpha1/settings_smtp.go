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

type SMTP struct {
	SMTPConfig  `json:",inline"`
	SMTPSecrets `json:",inline"`
}

type SMTPConfig struct {
	//	SMTP_IgnoreTLS
	//	SMTP_Pool
	Host      parameters.Parameter `json:"host,omitempty" env:"OVERWRITE_SETTING_SMTP_Host"`
	Port      parameters.Parameter `json:"port,omitempty" env:"OVERWRITE_SETTING_SMTP_Port"`
	FromEmail parameters.Parameter `json:"fromEmail,omitempty" env:"OVERWRITE_SETTING_From_Email"`
	Username  parameters.Parameter `json:"username,omitempty" env:"OVERWRITE_SETTING_SMTP_Username"`
}

type SMTPSecrets struct {
	Password parameters.Parameter `json:"passwor,omitempty" env:"OVERWRITE_SETTING_SMTP_Password"`
}

func (s *SMTP) SetDefaults() {
}

func (s *SMTP) GetParameters() parameters.Parameters {
	params, _ := parameters.Marshal(s.SMTPConfig)
	secretParams, _ := parameters.Marshal(s.SMTPSecrets)

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
