/*
Copyright 2021 Juicedata Inc

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

package common

// Runtime for Juice
const (
	JUICEFS_RUNTIME = "juicefs"

	JUICEFS_MOUNT_TYPE = "JuiceFS"

	JUICEFS_NAMESPACE = "juicefs-system"

	JUICEFS_CHART = JUICEFS_MOUNT_TYPE

	JUICEFS_FUSE_IMAGE_ENV = "JUICEFS_FUSE_IMAGE_ENV"

	DEFAULT_JUICEFS_FUSE_IMAGE = "juicedata/juicefs-csi-driver:v0.10.5"

	JuiceFSMountPath = "/bin/mount.juicefs"
)
