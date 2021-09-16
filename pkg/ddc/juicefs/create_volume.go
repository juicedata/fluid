package juicefs

import (
	"github.com/fluid-cloudnative/fluid/pkg/common"
	volumeHelper "github.com/fluid-cloudnative/fluid/pkg/utils/dataset/volume"
)

func (j JuiceFSEngine) CreateVolume() (err error) {
	if j.runtime == nil {
		j.runtime, err = j.getRuntime()
		if err != nil {
			return
		}
	}

	err = j.createFusePersistentVolume()
	if err != nil {
		return err
	}

	err = j.createFusePersistentVolumeClaim()
	if err != nil {
		return err
	}
	return
}

// createFusePersistentVolume
func (j *JuiceFSEngine) createFusePersistentVolume() (err error) {
	runtimeInfo, err := j.getRuntimeInfo()
	if err != nil {
		return err
	}

	return volumeHelper.CreatePersistentVolumeForRuntime(j.Client,
		runtimeInfo,
		j.getMountPoint(),
		common.JUICEFS_MOUNT_TYPE,
		j.Log)
}

// createFusePersistentVolume
func (j *JuiceFSEngine) createFusePersistentVolumeClaim() (err error) {
	runtimeInfo, err := j.getRuntimeInfo()
	if err != nil {
		return err
	}

	return volumeHelper.CreatePersistentVolumeClaimForRuntime(j.Client, runtimeInfo, j.Log)
}
