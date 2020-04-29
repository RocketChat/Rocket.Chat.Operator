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
	"fmt"

	instance "k8s.libre.sh/instance"
	meta "k8s.libre.sh/meta"
	parameters "k8s.libre.sh/parameters"
)

func (s *SMTP) SetDefaults() {
}

func (s *Database) SetDefaults() {
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
}

func (s *Settings) GetParameters() parameters.Parameters {

	params, err := parameters.Marshal(s.Database)

	for _, p := range params {
		if len(p.MountType) == 0 {
			p.MountType = "envFrom"
			if len(p.ValueFrom.Source) == 0 {
				if len(s.GetName()) > 0 {
					p.ValueFrom.Source = fmt.Sprintf("%v-%v", s.GetName(), s.GetComponent())
				}
			}
		}
	}
	if err != nil {

	}

	return params
}

func (s *Settings) GetObjects() map[int]instance.Object {

	return instance.GetObjects(s)
}
