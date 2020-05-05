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
	instance "k8s.libre.sh/instance"
	meta "k8s.libre.sh/meta"
	parameters "k8s.libre.sh/parameters"
)

func (s *SMTP) SetDefaults() {
}

func (s *SMTP) GetParameters() parameters.Parameters {
	params, _ := parameters.Marshal(s.SMTPConfig)
	secretParams, _ := parameters.Marshal(s.SMTPSecrets)

	params = append(params, secretParams...)

	return params
}

func (s *Database) SetDefaults() {
}

func (s *Database) GetParameters() parameters.Parameters {
	params, _ := parameters.Marshal(s.DatabaseConfig)
	secretParams, _ := parameters.Marshal(s.DatabaseSecrets)

	params = append(params, secretParams...)

	return params
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

	params = append(params, secretParams...)

	return params
}

func (s *Settings) GetMeta() instance.Meta { return s.ObjectMeta }

func (s *Settings) SetDefaults(i instance.Instance) {
	if s.ObjectMeta == nil {
		s.ObjectMeta = new(meta.ObjectMeta)
	}

	if len(s.ObjectMeta.Labels) == 0 {
		s.ObjectMeta.Labels = make(map[string]string)
	}

	s.ObjectMeta.Labels["app.kubernetes.io/component"] = "settings"

	instance.SetObjectMeta(i, s.ObjectMeta)

	s.General.SetDefaults()
}

func (s *Settings) GetParameters() parameters.Parameters {

	params := s.General.GetParameters()
	params = append(params, s.Database.GetParameters()...)
	params = append(params, s.SMTP.GetParameters()...)

	// Settings defaults for parameters
	for _, p := range params {
		if len(p.MountType) == 0 {
			p.MountType = parameters.MountFrom
			if len(p.ValueFrom.Ref) == 0 {
				if len(s.GetName()) > 0 {
					p.ValueFrom.Ref = s.GetName()
				}
			}
			if len(p.Type) == 0 {
				p.Type = parameters.ConfigParameter
			}
		}
	}

	return params
}

func (s *Settings) GetObjects() map[int]instance.Object {

	return instance.GetObjects(s)
}
