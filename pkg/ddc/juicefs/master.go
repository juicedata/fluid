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
	"context"
	datav1alpha1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
	"github.com/fluid-cloudnative/fluid/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/util/retry"
	"reflect"
)

func (j JuiceFSEngine) CheckMasterReady() (ready bool, err error) {
	return true, nil
}

func (j JuiceFSEngine) ShouldSetupMaster() (should bool, err error) {
	runtime, err := j.getRuntime()
	if err != nil {
		return
	}

	switch runtime.Status.FusePhase {
	case datav1alpha1.RuntimePhaseNone:
		should = true
	default:
		should = false
	}
	return
}

func (j JuiceFSEngine) SetupMaster() (err error) {
	fuseName := j.getFuseDaemonsetName()

	// 1. Setup the fuse
	_, err = j.getDaemonset(fuseName, j.namespace)
	if err != nil && apierrs.IsNotFound(err) {
		//1. Is not found error
		j.Log.V(1).Info("SetupMaster", "fuse", fuseName)
		return j.setupMasterInternal()
	} else if err != nil {
		//2. Other errors
		return
	} else {
		//3.The fuse has been set up
		j.Log.V(1).Info("The fuse has been set.")
	}

	// 2. Update the status of the runtime
	err = retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		runtime, err := j.getRuntime()
		if err != nil {
			return err
		}
		runtimeToUpdate := runtime.DeepCopy()

		runtimeToUpdate.Status.FusePhase = datav1alpha1.RuntimePhaseNotReady
		if len(runtimeToUpdate.Status.Conditions) == 0 {
			runtimeToUpdate.Status.Conditions = []datav1alpha1.RuntimeCondition{}
		}
		cond := utils.NewRuntimeCondition(datav1alpha1.RuntimeFusesInitialized, datav1alpha1.RuntimeFusesInitializedReason,
			"The fuse is initialized.", corev1.ConditionTrue)
		runtimeToUpdate.Status.Conditions =
			utils.UpdateRuntimeCondition(runtimeToUpdate.Status.Conditions,
				cond)

		if !reflect.DeepEqual(runtime.Status, runtimeToUpdate.Status) {
			return j.Client.Status().Update(context.TODO(), runtimeToUpdate)
		}

		return nil
	})

	if err != nil {
		j.Log.Error(err, "Update runtime status")
		return err
	}

	return
}
