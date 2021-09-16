package juicefs

import (
	"github.com/fluid-cloudnative/fluid/pkg/ddc/base"
	"github.com/fluid-cloudnative/fluid/pkg/utils"
	"github.com/fluid-cloudnative/fluid/pkg/utils/dataset/volume"
)

// getRuntimeInfo gets runtime info
func (j *JuiceFSEngine) getRuntimeInfo() (base.RuntimeInfoInterface, error) {
	if j.runtimeInfo == nil {
		runtime, err := j.getRuntime()
		if err != nil {
			return j.runtimeInfo, err
		}

		j.runtimeInfo, err = base.BuildRuntimeInfo(j.name, j.namespace, j.runtimeType, runtime.Spec.TieredStore)
		if err != nil {
			return j.runtimeInfo, err
		}

		// Setup Fuse Deploy Mode
		if runtime.Spec.Fuse.Global {
			j.runtimeInfo.SetupFuseDeployMode(runtime.Spec.Fuse.Global, runtime.Spec.Fuse.NodeSelector)
			j.Log.Info("Enable global mode for fuse")
		} else {
			j.Log.Info("Disable global mode for fuse")
		}

		if !j.UnitTest {
			// Check if the runtime is using deprecated labels
			isLabelDeprecated, err := j.HasDeprecatedCommonLabelName()
			if err != nil {
				return j.runtimeInfo, err
			}
			j.runtimeInfo.SetDeprecatedNodeLabel(isLabelDeprecated)

			// Check if the runtime is using deprecated naming style for PersistentVolumes
			isPVNameDeprecated, err := volume.HasDeprecatedPersistentVolumeName(j.Client, j.runtimeInfo, j.Log)
			if err != nil {
				return j.runtimeInfo, err
			}
			j.runtimeInfo.SetDeprecatedPVName(isPVNameDeprecated)

			j.Log.Info("Deprecation check finished", "isLabelDeprecated", j.runtimeInfo.IsDeprecatedNodeLabel(), "isPVNameDeprecated", j.runtimeInfo.IsDeprecatedPVName())

			// Setup with Dataset Info
			dataset, err := utils.GetDataset(j.Client, j.name, j.namespace)
			if err != nil {
				if utils.IgnoreNotFound(err) == nil {
					j.Log.Info("Dataset is notfound", "name", j.name, "namespace", j.namespace)
					return j.runtimeInfo, nil
				}

				j.Log.Info("Failed to get dataset when getruntimeInfo")
				return j.runtimeInfo, err
			}

			j.runtimeInfo.SetupWithDataset(dataset)

			j.Log.Info("Setup with dataset done", "exclusive", j.runtimeInfo.IsExclusive())
		}
	}

	return j.runtimeInfo, nil
}
