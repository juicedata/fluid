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
	"context"
	"fmt"
	datav1alpha1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
	"github.com/fluid-cloudnative/fluid/pkg/common"
	"github.com/fluid-cloudnative/fluid/pkg/utils"
	"github.com/fluid-cloudnative/fluid/pkg/utils/docker"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"strings"
)

// getRuntime gets the juicefs runtime
func (j *JuiceFSEngine) getRuntime() (*datav1alpha1.JuiceFSRuntime, error) {

	key := types.NamespacedName{
		Name:      j.name,
		Namespace: j.namespace,
	}

	var runtime datav1alpha1.JuiceFSRuntime
	if err := j.Get(context.TODO(), key, &runtime); err != nil {
		return nil, err
	}
	return &runtime, nil
}

func (j *JuiceFSEngine) getFuseDaemonsetName() (dsName string) {
	return j.name + "-fuse"
}

func (j *JuiceFSEngine) getDaemonset(name string, namespace string) (fuse *appsv1.DaemonSet, err error) {
	fuse = &appsv1.DaemonSet{}
	err = j.Client.Get(context.TODO(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, fuse)

	return fuse, err
}
func (j *JuiceFSEngine) getSecret(name string, namespace string) (fuse *corev1.Secret, err error) {
	fuse = &corev1.Secret{}
	err = j.Client.Get(context.TODO(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, fuse)

	return fuse, err
}

func (j *JuiceFSEngine) parseFuseImage(image string, tag string, imagePullPolicy string) (string, string, string) {
	if len(imagePullPolicy) == 0 {
		imagePullPolicy = common.DefaultImagePullPolicy
	}

	if len(image) == 0 {
		image = docker.GetImageRepoFromEnv(common.JUICEFS_FUSE_IMAGE_ENV)
		if len(image) == 0 {
			fuseImageInfo := strings.Split(common.DEFAULT_JUICEFS_FUSE_IMAGE, ":")
			if len(fuseImageInfo) < 1 {
				panic("invalid default juicefs fuse image!")
			} else {
				image = fuseImageInfo[0]
			}
		}
	}

	if len(tag) == 0 {
		tag = docker.GetImageTagFromEnv(common.JINDO_FUSE_IMAGE_ENV)
		if len(tag) == 0 {
			fuseImageInfo := strings.Split(common.DEFAULT_JUICEFS_FUSE_IMAGE, ":")
			if len(fuseImageInfo) < 2 {
				panic("invalid default init image!")
			} else {
				tag = fuseImageInfo[1]
			}
		}
	}

	return image, tag, imagePullPolicy
}

func (j *JuiceFSEngine) getMountPoint() (mountPath string) {
	mountRoot := getMountRoot()
	j.Log.Info("mountRoot", "path", mountRoot)
	return fmt.Sprintf("%s/%s/%s/juicefs-fuse", mountRoot, j.namespace, j.name)
}

// getMountRoot returns the default path, if it's not set
func getMountRoot() (path string) {
	path, err := utils.GetMountRoot()
	if err != nil {
		path = "/" + common.JUICEFS_RUNTIME
	} else {
		path = path + "/" + common.JUICEFS_RUNTIME
	}
	// e.Log.Info("Mount root", "path", path)
	return
}
