// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.libre.sh/application/components"
	"k8s.libre.sh/meta"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *App) DeepCopyInto(out *App) {
	*out = *in
	if in.Workload != nil {
		in, out := &in.Workload, &out.Workload
		*out = new(components.Workload)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new App.
func (in *App) DeepCopy() *App {
	if in == nil {
		return nil
	}
	out := new(App)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentList) DeepCopyInto(out *ComponentList) {
	*out = *in
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]ObjectStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentList.
func (in *ComponentList) DeepCopy() *ComponentList {
	if in == nil {
		return nil
	}
	out := new(ComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	out.DatabaseConfig = in.DatabaseConfig
	out.DatabaseSecrets = in.DatabaseSecrets
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseConfig) DeepCopyInto(out *DatabaseConfig) {
	*out = *in
	out.ReplicaSet = in.ReplicaSet
	out.AuthenticationDB = in.AuthenticationDB
	out.Username = in.Username
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseConfig.
func (in *DatabaseConfig) DeepCopy() *DatabaseConfig {
	if in == nil {
		return nil
	}
	out := new(DatabaseConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSecrets) DeepCopyInto(out *DatabaseSecrets) {
	*out = *in
	out.Password = in.Password
	out.URL = in.URL
	out.OplogURL = in.OplogURL
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSecrets.
func (in *DatabaseSecrets) DeepCopy() *DatabaseSecrets {
	if in == nil {
		return nil
	}
	out := new(DatabaseSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileUpload) DeepCopyInto(out *FileUpload) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileUpload.
func (in *FileUpload) DeepCopy() *FileUpload {
	if in == nil {
		return nil
	}
	out := new(FileUpload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *General) DeepCopyInto(out *General) {
	*out = *in
	out.GeneralConfig = in.GeneralConfig
	out.GeneralSecrets = in.GeneralSecrets
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new General.
func (in *General) DeepCopy() *General {
	if in == nil {
		return nil
	}
	out := new(General)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeneralConfig) DeepCopyInto(out *GeneralConfig) {
	*out = *in
	out.URL = in.URL
	out.InstanceIP = in.InstanceIP
	out.Username = in.Username
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeneralConfig.
func (in *GeneralConfig) DeepCopy() *GeneralConfig {
	if in == nil {
		return nil
	}
	out := new(GeneralConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GeneralSecrets) DeepCopyInto(out *GeneralSecrets) {
	*out = *in
	out.Password = in.Password
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GeneralSecrets.
func (in *GeneralSecrets) DeepCopy() *GeneralSecrets {
	if in == nil {
		return nil
	}
	out := new(GeneralSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStatus) DeepCopyInto(out *ObjectStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStatus.
func (in *ObjectStatus) DeepCopy() *ObjectStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStore) DeepCopyInto(out *ObjectStore) {
	*out = *in
	out.Bucket = in.Bucket
	out.Region = in.Region
	out.Endpoint = in.Endpoint
	out.PathStyle = in.PathStyle
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStore.
func (in *ObjectStore) DeepCopy() *ObjectStore {
	if in == nil {
		return nil
	}
	out := new(ObjectStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rocketchat) DeepCopyInto(out *Rocketchat) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rocketchat.
func (in *Rocketchat) DeepCopy() *Rocketchat {
	if in == nil {
		return nil
	}
	out := new(Rocketchat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Rocketchat) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RocketchatList) DeepCopyInto(out *RocketchatList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Rocketchat, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RocketchatList.
func (in *RocketchatList) DeepCopy() *RocketchatList {
	if in == nil {
		return nil
	}
	out := new(RocketchatList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RocketchatList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RocketchatSpec) DeepCopyInto(out *RocketchatSpec) {
	*out = *in
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = new(Settings)
		(*in).DeepCopyInto(*out)
	}
	out.Storage = in.Storage
	if in.App != nil {
		in, out := &in.App, &out.App
		*out = new(App)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RocketchatSpec.
func (in *RocketchatSpec) DeepCopy() *RocketchatSpec {
	if in == nil {
		return nil
	}
	out := new(RocketchatSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RocketchatStatus) DeepCopyInto(out *RocketchatStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ComponentList.DeepCopyInto(&out.ComponentList)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RocketchatStatus.
func (in *RocketchatStatus) DeepCopy() *RocketchatStatus {
	if in == nil {
		return nil
	}
	out := new(RocketchatStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SMTP) DeepCopyInto(out *SMTP) {
	*out = *in
	out.SMTPConfig = in.SMTPConfig
	out.SMTPSecrets = in.SMTPSecrets
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SMTP.
func (in *SMTP) DeepCopy() *SMTP {
	if in == nil {
		return nil
	}
	out := new(SMTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SMTPConfig) DeepCopyInto(out *SMTPConfig) {
	*out = *in
	out.Host = in.Host
	out.Port = in.Port
	out.FromEmail = in.FromEmail
	out.Username = in.Username
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SMTPConfig.
func (in *SMTPConfig) DeepCopy() *SMTPConfig {
	if in == nil {
		return nil
	}
	out := new(SMTPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SMTPSecrets) DeepCopyInto(out *SMTPSecrets) {
	*out = *in
	out.Password = in.Password
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SMTPSecrets.
func (in *SMTPSecrets) DeepCopy() *SMTPSecrets {
	if in == nil {
		return nil
	}
	out := new(SMTPSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Settings) DeepCopyInto(out *Settings) {
	*out = *in
	if in.ObjectMeta != nil {
		in, out := &in.ObjectMeta, &out.ObjectMeta
		*out = new(meta.ObjectMeta)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretMeta != nil {
		in, out := &in.SecretMeta, &out.SecretMeta
		*out = new(meta.ObjectMeta)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigMeta != nil {
		in, out := &in.ConfigMeta, &out.ConfigMeta
		*out = new(meta.ObjectMeta)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigRefs != nil {
		in, out := &in.ConfigRefs, &out.ConfigRefs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecretRefs != nil {
		in, out := &in.SecretRefs, &out.SecretRefs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Database = in.Database
	out.SMTP = in.SMTP
	out.General = in.General
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Settings.
func (in *Settings) DeepCopy() *Settings {
	if in == nil {
		return nil
	}
	out := new(Settings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	out.ObjectStore = in.ObjectStore
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}