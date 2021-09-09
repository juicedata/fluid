package juicefs

import (
	datav1alpha1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
	cruntime "github.com/fluid-cloudnative/fluid/pkg/runtime"
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

func (j JuiceFSEngine) CheckMasterReady() (ready bool, err error) {
	panic("implement me")
}

func (j JuiceFSEngine) CheckWorkersReady() (ready bool, err error) {
	panic("implement me")
}

func (j JuiceFSEngine) ShouldSetupMaster() (should bool, err error) {
	panic("implement me")
}

func (j JuiceFSEngine) ShouldSetupWorkers() (should bool, err error) {
	panic("implement me")
}

func (j JuiceFSEngine) ShouldCheckUFS() (should bool, err error) {
	panic("implement me")
}

func (j JuiceFSEngine) SetupMaster() (err error) {
	panic("implement me")
}

func (j JuiceFSEngine) SetupWorkers() (err error) {
	panic("implement me")
}

func (j JuiceFSEngine) UpdateDatasetStatus(phase datav1alpha1.DatasetPhase) (err error) {
	panic("implement me")
}

func (j JuiceFSEngine) PrepareUFS() (err error) {
	panic("implement me")
}

func (j JuiceFSEngine) ShouldUpdateUFS() (ufsToUpdate *utils.UFSToUpdate) {
	panic("implement me")
}

func (j JuiceFSEngine) UpdateOnUFSChange(ufsToUpdate *utils.UFSToUpdate) (ready bool, err error) {
	panic("implement me")
}

func (j JuiceFSEngine) Shutdown() error {
	panic("implement me")
}

func (j JuiceFSEngine) AssignNodesToCache(desiredNum int32) (currentNum int32, err error) {
	panic("implement me")
}

func (j JuiceFSEngine) CheckRuntimeHealthy() (err error) {
	panic("implement me")
}

func (j JuiceFSEngine) UpdateCacheOfDataset() (err error) {
	panic("implement me")
}

func (j JuiceFSEngine) CheckAndUpdateRuntimeStatus() (ready bool, err error) {
	panic("implement me")
}

func (j JuiceFSEngine) CreateVolume() error {
	panic("implement me")
}

func (j JuiceFSEngine) SyncReplicas(ctx cruntime.ReconcileRequestContext) error {
	panic("implement me")
}

func (j JuiceFSEngine) SyncMetadata() (err error) {
	panic("implement me")
}

func (j JuiceFSEngine) DeleteVolume() (err error) {
	panic("implement me")
}

func (j JuiceFSEngine) BindToDataset() (err error) {
	panic("implement me")
}

func (j JuiceFSEngine) CreateDataLoadJob(ctx cruntime.ReconcileRequestContext, targetDataload datav1alpha1.DataLoad) error {
	panic("implement me")
}

func (j JuiceFSEngine) CheckRuntimeReady() (ready bool) {
	panic("implement me")
}

func (j JuiceFSEngine) CheckExistenceOfPath(targetDataload datav1alpha1.DataLoad) (notExist bool, err error) {
	panic("implement me")
}
