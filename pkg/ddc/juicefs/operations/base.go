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

package operations

import (
	"context"
	"fmt"
	"github.com/fluid-cloudnative/fluid/pkg/utils/kubeclient"
	"github.com/go-logr/logr"
	"strings"
	"time"
)

type JuiceFileUtils struct {
	podName   string
	namespace string
	container string
	log       logr.Logger
}

func NewJuiceFileUtils(podName string, containerName string, namespace string, log logr.Logger) JuiceFileUtils {
	return JuiceFileUtils{
		podName:   podName,
		namespace: namespace,
		container: containerName,
		log:       log,
	}
}

// IsExist checks if the juicePath exists
func (j JuiceFileUtils) IsExist(juiceSubPath string) (found bool, err error) {
	var (
		command = []string{"ls", juiceSubPath}
		stdout  string
		stderr  string
	)

	stdout, stderr, err = j.exec(command, true)
	if err != nil {
		if strings.Contains(stdout, "No such file or directory") || strings.Contains(stderr, "No such file or directory") {
			return false, nil
		} else {
			err = fmt.Errorf("execute command %v with expectedErr: %v stdout %s and stderr %s", command, err, stdout, stderr)
			return false, err
		}
	} else {
		found = true
	}
	return
}

// Mkdir mkdir in juicefs container
func (j JuiceFileUtils) Mkdir(juiceSubPath string) (err error) {
	var (
		command = []string{"mkdir", juiceSubPath}
		stdout  string
		stderr  string
	)

	stdout, stderr, err = j.exec(command, true)
	if err != nil {
		if strings.Contains(stdout, "File exists") {
			err = nil
		} else {
			err = fmt.Errorf("execute command %v with expectedErr: %v stdout %s and stderr %s", command, err, stdout, stderr)
			return
		}
	}
	return
}

// DeleteDir delete dir in pod
func (j JuiceFileUtils) DeleteDir(dir string) (err error) {
	var (
		command = []string{"rm", "-rf", dir}
		stdout  string
		stderr  string
	)

	stdout, stderr, err = j.exec(command, true)
	if err != nil {
		err = fmt.Errorf("execute command %v with expectedErr: %v stdout %s and stderr %s", command, err, stdout, stderr)
		return
	}
	return
}

// GetMetric Get pod metrics
func (j JuiceFileUtils) GetMetric() (metrics string, err error) {
	var (
		command = []string{"curl", "0.0.0.0:9567/metrics"}
		stdout  string
		stderr  string
	)

	stdout, stderr, err = j.exec(command, true)
	if err != nil {
		err = fmt.Errorf("execute command %v with expectedErr: %v stdout %s and stderr %s", command, err, stdout, stderr)
		return
	}
	metrics = stdout
	return
}

// exec with timeout
func (j JuiceFileUtils) exec(command []string, verbose bool) (stdout string, stderr string, err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1500)
	ch := make(chan string, 1)
	defer cancel()

	go func() {
		stdout, stderr, err = j.execWithoutTimeout(command, verbose)
		ch <- "done"
	}()

	select {
	case <-ch:
		j.log.Info("execute in time", "command", command)
	case <-ctx.Done():
		err = fmt.Errorf("timeout when executing %v", command)
	}
	return
}

// execWithoutTimeout
func (j JuiceFileUtils) execWithoutTimeout(command []string, verbose bool) (stdout string, stderr string, err error) {
	stdout, stderr, err = kubeclient.ExecCommandInContainer(j.podName, j.container, j.namespace, command)
	if err != nil {
		j.log.Info("Stdout", "Command", command, "Stdout", stdout)
		j.log.Error(err, "Failed", "Command", command, "FailedReason", stderr)
		return
	}
	if verbose {
		j.log.Info("Stdout", "Command", command, "Stdout", stdout)
	}

	return
}
