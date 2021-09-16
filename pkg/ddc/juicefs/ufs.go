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
