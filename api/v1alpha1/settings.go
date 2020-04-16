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
	meta "k8s.libre.sh/meta"
	parameters "k8s.libre.sh/parameters"
)

type SMTP struct {
	Host      parameters.Parameter `json:"host,omitempty" env:"OVERWRITE_SETTING_SMTP_Host"`
	Port      parameters.Parameter `json:"port,omitempty" env:"OVERWRITE_SETTING_SMTP_Port"`
	Username  parameters.Parameter `json:"username,omitempty" env:"OVERWRITE_SETTING_SMTP_Username"`
	FromEmail parameters.Parameter `json:"fromEmail,omitempty" env:"OVERWRITE_SETTING_From_Email"`
	Password  parameters.Parameter `json:"passwor,omitempty" env:"OVERWRITE_SETTING_SMTP_Password"`
}

type Database struct {
	URL      parameters.Parameter `json:"url,omitempty" env:"MONGO_URL"`
	OplogURL parameters.Parameter `json:"oplogURL,omitempty" env:"MONGO_OPLOG_URL"`
}

type Settings struct {
	*meta.ObjectMeta `json:"meta,omitempty"`
	Database         Database `json:"database,omitempty"`
	SMTP             SMTP     `json:"smtp,omitempty"`
}
