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

type Database struct {
	DatabaseConfig  `json:",inline"`
	DatabaseSecrets `json:",inline"`
}

type DatabaseConfig struct {
	ReplicaSet       parameters.Parameter `json:"replicaSet,omitempty" env:"MONGO_REPLICASET"`
	AuthenticationDB parameters.Parameter `json:"authenticationDB,omitempty" env:"MONGO_AUTHDB"`
	Username         parameters.Parameter `json:"password,omitempty" env:"MONGO_USERNAME"`
}

type DatabaseSecrets struct {
	Password parameters.Parameter `json:"username,omitempty" env:"MONGO_PASSWORD"`
	URL      parameters.Parameter `json:"url,omitempty" env:"MONGO_URL"`
	OplogURL parameters.Parameter `json:"oplogURL,omitempty" env:"MONGO_OPLOG_URL"`
}

func (s *Database) SetDefaults() {
}

func (s *Database) GetParameters() parameters.Parameters {
	params, _ := parameters.Marshal(s.DatabaseConfig)
	secretParams, _ := parameters.Marshal(s.DatabaseSecrets)

	for _, p := range secretParams {
		if len(p.Type) == 0 {
			p.Type = parameters.SecretParameter
		}
		if len(p.ValueFrom.Ref) == 0 && len(p.Generate) == 0 && len(p.Value) == 0 {
			p.Generate = parameters.GenerateRand12
		}
		if len(p.MountType) == 0 {
			p.MountType = parameters.MountEnvFile
		}
	}

	params = append(params, secretParams...)

	return params
}
