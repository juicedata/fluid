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
	value.Fuse.HostMountPath = mount.MountPoint
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
