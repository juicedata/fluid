package juicefs

import volumeHelper "github.com/fluid-cloudnative/fluid/pkg/utils/dataset/volume"

func (j JuiceFSEngine) DeleteVolume() (err error) {
	if j.runtime == nil {
		j.runtime, err = j.getRuntime()
		if err != nil {
			return
		}
	}

	err = j.deleteFusePersistentVolumeClaim()
	if err != nil {
		return
	}

	err = j.deleteFusePersistentVolume()
	if err != nil {
		return
	}

	return
}

// deleteFusePersistentVolume
func (j *JuiceFSEngine) deleteFusePersistentVolume() (err error) {
	runtimeInfo, err := j.getRuntimeInfo()
	if err != nil {
		return err
	}

	return volumeHelper.DeleteFusePersistentVolume(j.Client, runtimeInfo, j.Log)
}

// deleteFusePersistentVolume
func (j *JuiceFSEngine) deleteFusePersistentVolumeClaim() (err error) {
	runtimeInfo, err := j.getRuntimeInfo()
	if err != nil {
		return err
	}

	return volumeHelper.DeleteFusePersistentVolumeClaim(j.Client, runtimeInfo, j.Log)
}
