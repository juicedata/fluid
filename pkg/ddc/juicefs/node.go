package juicefs

import (
	"github.com/fluid-cloudnative/fluid/pkg/utils"
	datasetSchedule "github.com/fluid-cloudnative/fluid/pkg/utils/dataset/lifecycle"
)

func (j JuiceFSEngine) AssignNodesToCache(desiredNum int32) (currentScheduleNum int32, err error) {
	runtimeInfo, err := j.getRuntimeInfo()
	if err != nil {
		return currentScheduleNum, err
	}

	dataset, err := utils.GetDataset(j.Client, j.name, j.namespace)
	if err != nil {
		return
	}

	j.Log.Info("AssignNodesToCache", "dataset", dataset)
	return datasetSchedule.AssignDatasetToNodes(runtimeInfo,
		dataset,
		j.Client,
		desiredNum)
}
