package juicefs

import "github.com/fluid-cloudnative/fluid/pkg/common"

func (j *JuiceFSEngine) getCommonLabelName() string {
	return common.LabelAnnotationStorageCapacityPrefix + j.namespace + "-" + j.name
}
