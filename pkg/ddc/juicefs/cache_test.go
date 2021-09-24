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
	"github.com/brahma-adshonor/gohook"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestJuiceFSEngine_queryCacheStatus(t *testing.T) {
	ReturnOnePods := func(a JuiceFSEngine, dsName string, namespace string) (pods []corev1.Pod, err error) {
		return []corev1.Pod{
			{ObjectMeta: metav1.ObjectMeta{Name: "test1"}},
		}, nil
	}
	PodMetrics := func(a JuiceFSEngine, podName string) (metrics string, err error) {
		return mockJuiceFSMetric(), nil
	}
	wrappedUnhookPods := func() {
		err := gohook.UnHook(JuiceFSEngine.getRunningPodsOfDaemonset)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	wrappedUnhookMetrics := func() {
		err := gohook.UnHook(JuiceFSEngine.getPodMetrics)
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	err := gohook.Hook(JuiceFSEngine.getRunningPodsOfDaemonset, ReturnOnePods, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = gohook.Hook(JuiceFSEngine.getPodMetrics, PodMetrics, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	a := JuiceFSEngine{
		name:        "test",
		namespace:   "default",
		runtimeType: "JuiceFSRuntime",
		Log:         nil,
	}
	want := cacheStates{
		cacheCapacity:        "",
		cached:               "37.81MiB",
		cachedPercentage:     "151.2%",
		cacheHitRatio:        "100.0%",
		cacheThroughputRatio: "100.0%",
	}
	got, err := a.queryCacheStatus()
	if err != nil {
		t.Error("check failure, want err, got nil")
	}
	if want != got {
		t.Errorf("got=%v, want=%v", got, want)
	}
	wrappedUnhookPods()
	wrappedUnhookMetrics()
}
