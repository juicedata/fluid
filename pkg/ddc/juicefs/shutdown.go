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
	"github.com/fluid-cloudnative/fluid/pkg/utils/helm"
)

func (j JuiceFSEngine) Shutdown() (err error) {
	if j.retryShutdown < j.gracefulShutdownLimits {
		err = j.cleanupCache()
		if err != nil {
			j.retryShutdown = j.retryShutdown + 1
			j.Log.Info("clean cache failed",
				"retry times", j.retryShutdown)
			return
		}
	}
	err = j.destroyMaster()
	if err != nil {
		return
	}
	return nil
}

// destroyMaster Destroy the master
func (j *JuiceFSEngine) destroyMaster() (err error) {
	var found bool
	found, err = helm.CheckRelease(j.name, j.namespace)
	if err != nil {
		return err
	}

	if found {
		err = helm.DeleteRelease(j.name, j.namespace)
		if err != nil {
			return
		}
	}
	return
}

// cleanupCache cleans up the cache
func (j *JuiceFSEngine) cleanupCache() (err error) {
	// todo
	return
}
