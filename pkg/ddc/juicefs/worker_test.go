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
	datav1alpha1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
	"github.com/fluid-cloudnative/fluid/pkg/ddc/base"
	"github.com/fluid-cloudnative/fluid/pkg/utils/kubeclient"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"testing"
)

func TestJuiceFSEngine_ShouldSetupWorkers(t *testing.T) {
	type fields struct {
		name      string
		namespace string
		runtime   *datav1alpha1.JuiceFSRuntime
	}
	tests := []struct {
		name       string
		fields     fields
		wantShould bool
		wantErr    bool
	}{
		{
			name: "test0",
			fields: fields{
				name:      "test0",
				namespace: "juicefs",
				runtime: &datav1alpha1.JuiceFSRuntime{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test0",
						Namespace: "juicefs",
					},
					Status: datav1alpha1.JuiceFSRuntimeStatus{
						FusePhase: datav1alpha1.RuntimePhaseNone,
					},
				},
			},
			wantShould: true,
			wantErr:    false,
		},
		{
			name: "test1",
			fields: fields{
				name:      "test1",
				namespace: "juicefs",
				runtime: &datav1alpha1.JuiceFSRuntime{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test1",
						Namespace: "juicefs",
					},
					Status: datav1alpha1.JuiceFSRuntimeStatus{
						FusePhase: datav1alpha1.RuntimePhaseNotReady,
					},
				},
			},
			wantShould: false,
			wantErr:    false,
		},
		{
			name: "test2",
			fields: fields{
				name:      "test2",
				namespace: "juicefs",
				runtime: &datav1alpha1.JuiceFSRuntime{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test2",
						Namespace: "juicefs",
					},
					Status: datav1alpha1.JuiceFSRuntimeStatus{
						FusePhase: datav1alpha1.RuntimePhasePartialReady,
					},
				},
			},
			wantShould: false,
			wantErr:    false,
		},
		{
			name: "test3",
			fields: fields{
				name:      "test3",
				namespace: "juicefs",
				runtime: &datav1alpha1.JuiceFSRuntime{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test3",
						Namespace: "juicefs",
					},
					Status: datav1alpha1.JuiceFSRuntimeStatus{
						FusePhase: datav1alpha1.RuntimePhaseReady,
					},
				},
			},
			wantShould: false,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runtimeObjs := []runtime.Object{}
			data := &datav1alpha1.Dataset{
				ObjectMeta: metav1.ObjectMeta{
					Name:      tt.fields.name,
					Namespace: tt.fields.namespace,
				},
			}

			s := runtime.NewScheme()
			s.AddKnownTypes(datav1alpha1.GroupVersion, tt.fields.runtime)
			s.AddKnownTypes(datav1alpha1.GroupVersion, data)
			_ = v1.AddToScheme(s)
			runtimeObjs = append(runtimeObjs, tt.fields.runtime, data)
			mockClient := fake.NewFakeClientWithScheme(s, runtimeObjs...)
			e := &JuiceFSEngine{
				name:      tt.fields.name,
				namespace: tt.fields.namespace,
				runtime:   tt.fields.runtime,
				Client:    mockClient,
			}

			gotShould, err := e.ShouldSetupWorkers()
			if (err != nil) != tt.wantErr {
				t.Errorf("JuiceFSEngine.ShouldSetupWorkers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotShould != tt.wantShould {
				t.Errorf("JuiceFSEngine.ShouldSetupWorkers() = %v, want %v", gotShould, tt.wantShould)
			}
		})
	}
}

func TestJuiceFSEngine_SetupWorkers(t *testing.T) {
	runtimeInfo, err := base.BuildRuntimeInfo("juicefs", "fluid", "juicefs", datav1alpha1.TieredStore{})

	if err != nil {
		t.Errorf("fail to create the runtimeInfo with error %v", err)
	}
	runtimeInfo.SetupWithDataset(&datav1alpha1.Dataset{
		Spec: datav1alpha1.DatasetSpec{PlacementMode: datav1alpha1.ExclusiveMode},
	})

	nodeSelector := map[string]string{
		"node-select": "true",
	}
	runtimeInfo.SetupFuseDeployMode(true, nodeSelector)

	type fields struct {
		replicas    int32
		nodeInputs  []*v1.Node
		runtime     *datav1alpha1.JuiceFSRuntime
		runtimeInfo base.RuntimeInfoInterface
		name        string
		namespace   string
	}
	tests := []struct {
		name             string
		fields           fields
		wantedNodeLabels map[string]map[string]string
	}{
		{
			name: "test0",
			fields: fields{
				replicas: 1,
				nodeInputs: []*v1.Node{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test-node",
						},
					},
				},
				runtime: &datav1alpha1.JuiceFSRuntime{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test",
						Namespace: "fluid",
					},
					Spec: datav1alpha1.JuiceFSRuntimeSpec{
						Replicas: 1,
					},
				},
				runtimeInfo: runtimeInfo,
				name:        "test",
				namespace:   "fluid",
			},
			wantedNodeLabels: map[string]map[string]string{
				"test-node": {
					"fluid.io/dataset-num":                 "1",
					"fluid.io/s-fluid-juicefs":             "true",
					"fluid.io/s-h-juicefs-t-fluid-juicefs": "0B",
					"fluid.io/s-juicefs-fluid-juicefs":     "true",
					"fluid_exclusive":                      "fluid_juicefs",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runtimeObjs := []runtime.Object{}
			for _, nodeInput := range tt.fields.nodeInputs {
				runtimeObjs = append(runtimeObjs, nodeInput.DeepCopy())
			}

			s := runtime.NewScheme()
			data := &datav1alpha1.Dataset{
				ObjectMeta: metav1.ObjectMeta{
					Name:      tt.fields.name,
					Namespace: tt.fields.namespace,
				},
			}
			s.AddKnownTypes(datav1alpha1.GroupVersion, tt.fields.runtime)
			s.AddKnownTypes(datav1alpha1.GroupVersion, data)
			_ = v1.AddToScheme(s)
			runtimeObjs = append(runtimeObjs, tt.fields.runtime)
			runtimeObjs = append(runtimeObjs, data)
			mockClient := fake.NewFakeClientWithScheme(s, runtimeObjs...)

			e := &JuiceFSEngine{
				runtime:     tt.fields.runtime,
				runtimeInfo: tt.fields.runtimeInfo,
				Client:      mockClient,
				name:        tt.fields.name,
				namespace:   tt.fields.namespace,
				Log:         ctrl.Log.WithName(tt.fields.name),
			}
			err := e.SetupWorkers()
			if err != nil {
				t.Errorf("JuiceFSEngine.SetupWorkers() error = %v", err)
			}
			for _, node := range tt.fields.nodeInputs {
				newNode, err := kubeclient.GetNode(mockClient, node.Name)
				if err != nil {
					t.Errorf("fail to get the node with the error %v", err)
				}

				if len(newNode.Labels) != len(tt.wantedNodeLabels[node.Name]) {
					t.Errorf("fail to decrease the labels, newNode labels is %v", newNode.Labels)
				}
				if len(newNode.Labels) != 0 && !reflect.DeepEqual(newNode.Labels, tt.wantedNodeLabels[node.Name]) {
					t.Errorf("fail to decrease the labels, newNode labels is %v", newNode.Labels)
				}
			}
		})
	}
}

func TestJuiceFSEngine_CheckWorkersReady(t *testing.T) {
	type fields struct {
		runtime   *datav1alpha1.JuiceFSRuntime
		fuse      *appsv1.DaemonSet
		name      string
		namespace string
	}
	tests := []struct {
		name      string
		fields    fields
		wantReady bool
		wantErr   bool
	}{
		{
			name: "test0",
			fields: fields{
				name:      "test0",
				namespace: "juicefs",
				runtime: &datav1alpha1.JuiceFSRuntime{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test0",
						Namespace: "juicefs",
					},
					Spec: datav1alpha1.JuiceFSRuntimeSpec{
						Replicas: 1,
						Fuse: datav1alpha1.JuiceFSFuseSpec{
							Global: true,
						},
					},
				},
				fuse: &appsv1.DaemonSet{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test0-fuse",
						Namespace: "juicefs",
					},
					Status: appsv1.DaemonSetStatus{
						NumberAvailable:        1,
						DesiredNumberScheduled: 1,
						CurrentNumberScheduled: 1,
					},
				},
			},
			wantReady: true,
			wantErr:   false,
		},
		{
			name: "test1",
			fields: fields{
				name:      "test1",
				namespace: "juicefs",
				runtime: &datav1alpha1.JuiceFSRuntime{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test1",
						Namespace: "juicefs",
					},
					Spec: datav1alpha1.JuiceFSRuntimeSpec{
						Replicas: 1,
						Fuse: datav1alpha1.JuiceFSFuseSpec{
							Global: true,
						},
					},
				},
				fuse: &appsv1.DaemonSet{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "test1-fuse",
						Namespace: "juicefs",
					},
					Status: appsv1.DaemonSetStatus{
						NumberAvailable:        0,
						DesiredNumberScheduled: 1,
						CurrentNumberScheduled: 0,
					},
				},
			},
			wantReady: false,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runtimeObjs := []runtime.Object{}
			data := &datav1alpha1.Dataset{
				ObjectMeta: metav1.ObjectMeta{
					Name:      tt.fields.name,
					Namespace: tt.fields.namespace,
				},
			}

			s := runtime.NewScheme()
			s.AddKnownTypes(datav1alpha1.GroupVersion, tt.fields.runtime)
			s.AddKnownTypes(datav1alpha1.GroupVersion, data)
			s.AddKnownTypes(appsv1.SchemeGroupVersion, tt.fields.fuse)
			_ = v1.AddToScheme(s)

			runtimeObjs = append(runtimeObjs, tt.fields.runtime, data, tt.fields.fuse)
			mockClient := fake.NewFakeClientWithScheme(s, runtimeObjs...)
			e := &JuiceFSEngine{
				runtime:   tt.fields.runtime,
				name:      tt.fields.name,
				namespace: tt.fields.namespace,
				Client:    mockClient,
				Log:       ctrl.Log.WithName(tt.fields.name),
			}
			gotReady, err := e.CheckWorkersReady()
			if (err != nil) != tt.wantErr {
				t.Errorf("JuiceFSEngine.CheckWorkersReady() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotReady != tt.wantReady {
				t.Errorf("JuiceFSEngine.CheckWorkersReady() = %v, want %v", gotReady, tt.wantReady)
			}
		})
	}
}