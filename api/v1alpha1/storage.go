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
	parameters "k8s.libre.sh/parameters"
)

type FileUpload struct {
	StorageType string `json:"storageType,omitempty" env:"OVERWRITE_SETTING_FileUpload_Storage_Type"`
}

// FileUpload_FileSystemPath

//FileUpload_Webdav_Upload_Folder_Path
//FileUpload_Webdav_Server_URL
//FileUpload_Webdav_Username
//FileUpload_Webdav_Password
//FileUpload_Webdav_Proxy_Avatars
//FileUpload_Webdav_Proxy_Uploads

type ObjectStore struct {
	//	FileUpload_S3_Acl
	//	FileUpload_S3_CDN
	//	FileUpload_S3_BucketURL
	//	FileUpload_S3_SignatureVersion
	//	FileUpload_S3_URLExpiryTimeSpan
	//	FileUpload_S3_Proxy_Avatars
	//	FileUpload_S3_Proxy_Uploads
	Bucket    parameters.Parameter `json:"bucket,omitempty" env:"OVERWRITE_SETTING_FileUpload_S3_Bucket"`
	Region    parameters.Parameter `json:"region,omitempty" env:"OVERWRITE_SETTING_FileUpload_S3_Region"`
	Endpoint  parameters.Parameter `json:"endpoint,omitempty" env:"OVERWRITE_SETTING_FileUpload_S3_BucketURL"`
	PathStyle parameters.Parameter `json:"pathStyle,omitempty" env:"OVERWRITE_SETTING_FileUpload_S3_ForcePathStyle"`
}

type Storage struct {
	ObjectStore ObjectStore `json:"objectStore,omitempty"`
}
