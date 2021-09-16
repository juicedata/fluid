package juicefs

import (
	"fmt"
	datav1alpha1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
	"github.com/fluid-cloudnative/fluid/pkg/utils"
)

func (j *JuiceFSEngine) transform(runtime *datav1alpha1.JuiceFSRuntime) (value *JuiceFS, err error) {
	if runtime == nil {
		err = fmt.Errorf("the juicefsRuntime is null")
		return
	}

	dataset, err := utils.GetDataset(j.Client, j.name, j.namespace)
	if err != nil {
		return value, err
	}

	value = &JuiceFS{}

	value.FullnameOverride = j.name

	// transform the fuse
	err = j.transformFuse(runtime, dataset, value)
	if err != nil {
		return
	}

	return
}
