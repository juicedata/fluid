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

import (
	"errors"
	"fmt"
	datav1alpha1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
	"github.com/fluid-cloudnative/fluid/pkg/common"
	"strings"
)

func (j *JuiceFSEngine) transformFuse(runtime *datav1alpha1.JuiceFSRuntime, dataset *datav1alpha1.Dataset, value *JuiceFS) (err error) {
	value.Fuse = Fuse{}

	if len(dataset.Spec.Mounts) <= 0 {
		return errors.New("do not assign mount point")
	}
	mount := dataset.Spec.Mounts[0]

	var secretName string
	if runtime.Spec.Fuse.SecretName == "" {
		// if runtime secretName is nil, use the same name as runtime
		secretName = runtime.Name
	} else {
		secretName = runtime.Spec.Fuse.SecretName
	}
	secret, err := j.getSecret(secretName, j.namespace)
	if err != nil {
		return
	}

	image := runtime.Spec.Fuse.Image
	tag := runtime.Spec.Fuse.ImageTag
	imagePullPolicy := runtime.Spec.Fuse.ImagePullPolicy

	value.Fuse.Image, value.Fuse.ImageTag, value.ImagePullPolicy = j.parseFuseImage(image, tag, imagePullPolicy)
	value.Fuse.MountPath = j.getMountPoint()
	value.Fuse.NodeSelector = map[string]string{}
	if strings.HasPrefix(mount.MountPoint, "local://") {
		value.Fuse.HostMountPath = mount.MountPoint[8:]
	} else {
		value.Fuse.HostMountPath = mount.MountPoint
	}
	if mount.Path == "" {
		value.Fuse.SubPath = mount.Name
	} else {
		value.Fuse.SubPath = mount.Path
	}

	mountArgs := []string{common.JuiceFSMountPath, string(secret.Data["name"]), value.Fuse.MountPath}
	options := []string{"metrics=0.0.0.0:9567"}
	for k, v := range mount.Options {
		options = append(options, fmt.Sprintf("%s=%s", k, v))
	}
	if len(runtime.Spec.TieredStore.Levels) >= 0 {
		cacheDir := runtime.Spec.TieredStore.Levels[0].Path
		cacheSize := runtime.Spec.TieredStore.Levels[0].Quota
		cacheRatio := runtime.Spec.TieredStore.Levels[0].Low
		options = append(options, fmt.Sprintf("cache-dir=%s", cacheDir))
		options = append(options, fmt.Sprintf("cache-size=%s", cacheSize))
		options = append(options, fmt.Sprintf("free-space-ratio=%s", cacheRatio))
	}

	mountArgs = append(mountArgs, "-o", strings.Join(options, ","))

	value.Fuse.Command = strings.Join(mountArgs, " ")
	value.Fuse.StatCmd = "stat -c %i " + value.Fuse.MountPath

	if runtime.Spec.Fuse.Global {
		if len(runtime.Spec.Fuse.NodeSelector) > 0 {
			value.Fuse.NodeSelector = runtime.Spec.Fuse.NodeSelector
		}
		value.Fuse.NodeSelector[common.FLUID_FUSE_BALLOON_KEY] = common.FLUID_FUSE_BALLOON_VALUE
		j.Log.Info("Enable Fuse's global mode")
	} else {
		labelName := j.getCommonLabelName()
		value.Fuse.NodeSelector[labelName] = "true"
		j.Log.Info("Disable Fuse's global mode")
	}

	value.Fuse.Enabled = true

	j.transformResourcesForFuse(runtime, value)

	return
}
