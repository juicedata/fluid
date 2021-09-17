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
	"github.com/fluid-cloudnative/fluid/pkg/utils"
)

func (j JuiceFSEngine) UsedStorageBytes() (int64, error) {
	panic("implement me")
}

func (j JuiceFSEngine) FreeStorageBytes() (int64, error) {
	panic("implement me")
}

func (j JuiceFSEngine) TotalStorageBytes() (int64, error) {
	panic("implement me")
}

func (j JuiceFSEngine) TotalFileNums() (int64, error) {
	panic("implement me")
}

func (j JuiceFSEngine) ShouldCheckUFS() (should bool, err error) {
	return false, nil
}

func (j JuiceFSEngine) PrepareUFS() (err error) {
	panic("implement me")
}

func (j JuiceFSEngine) ShouldUpdateUFS() (ufsToUpdate *utils.UFSToUpdate) {
	return nil
}

func (j JuiceFSEngine) UpdateOnUFSChange(ufsToUpdate *utils.UFSToUpdate) (ready bool, err error) {
	panic("implement me")
}
