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

package juicefs

import "github.com/fluid-cloudnative/fluid/pkg/common"

// JuiceFS The value yaml file
type JuiceFS struct {
	FullnameOverride string `yaml:"fullnameOverride"`

	common.ImageInfo `yaml:",inline"`
	common.UserInfo  `yaml:",inline"`

	NodeSelector map[string]string `yaml:"nodeSelector,omitempty"`
	Fuse         Fuse              `yaml:"fuse,omitempty"`
	TieredStore  TieredStore       `yaml:"tieredstore,omitempty"`
}

type Fuse struct {
	Image           string            `yaml:"image,omitempty"`
	NodeSelector    map[string]string `yaml:"nodeSelector,omitempty"`
	ImageTag        string            `yaml:"imageTag,omitempty"`
	ImagePullPolicy string            `yaml:"imagePullPolicy,omitempty"`
	MountPath       string            `yaml:"mountPath,omitempty"`
	SubPath         string            `yaml:"subPath,omitempty"`
	HostMountPath   string            `yaml:"hostMountPath,omitempty"`
	Command         string            `yaml:"command,omitempty"`
	StatCmd         string            `yaml:"statCmd,omitempty"`
	Enabled         bool              `yaml:"enabled,omitempty"`
	Resources       common.Resources  `yaml:"resources,omitempty"`
}

type TieredStore struct {
	Path  string `yaml:"path,omitempty"`
	Quota string `yaml:"quota,omitempty"`
	High  string `yaml:"high,omitempty"`
	Low   string `yaml:"low,omitempty"`
}
