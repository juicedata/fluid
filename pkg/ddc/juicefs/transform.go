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
