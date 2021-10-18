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
	"testing"

	"github.com/brahma-adshonor/gohook"
	datav1alpha1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
	"github.com/fluid-cloudnative/fluid/pkg/ddc/juicefs/operations"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func TestShouldSyncMetadata(t *testing.T) {
	datasetInputs := []datav1alpha1.Dataset{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "hbase",
				Namespace: "fluid",
			},
			Status: datav1alpha1.DatasetStatus{
				UfsTotal: "2Gi",
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "spark",
				Namespace: "fluid",
			},
			Status: datav1alpha1.DatasetStatus{
				UfsTotal: "",
			},
		},
	}
	testObjs := []runtime.Object{}
	for _, datasetInput := range datasetInputs {
		testObjs = append(testObjs, datasetInput.DeepCopy())
	}
	client := fake.NewFakeClientWithScheme(testScheme, testObjs...)

	engines := []JuiceFSEngine{
		{
			name:      "hbase",
			namespace: "fluid",
			Client:    client,
			Log:       log.NullLogger{},
		},
		{
			name:      "spark",
			namespace: "fluid",
			Client:    client,
			Log:       log.NullLogger{},
		},
	}

	var testCases = []struct {
		engine         JuiceFSEngine
		expectedShould bool
	}{
		{
			engine:         engines[0],
			expectedShould: false,
		},
		{
			engine:         engines[1],
			expectedShould: true,
		},
	}

	for _, test := range testCases {
		should, err := test.engine.shouldSyncMetadata()
		if err != nil || should != test.expectedShould {
			t.Errorf("fail to exec the function")
		}
	}
}

func TestSyncMetadata(t *testing.T) {
	QueryMetaDataInfoIntoFileCommon := func(a operations.JuiceFileUtils, key operations.KeyOfMetaDataFile, filename string) (value string, err error) {
		return "1024", nil
	}

	wrappedUnhookQueryMetaDataInfoIntoFile := func() {
		err := gohook.UnHook(operations.JuiceFileUtils.QueryMetaDataInfoIntoFile)
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	datasetInputs := []datav1alpha1.Dataset{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "hbase",
				Namespace: "fluid",
			},
			Status: datav1alpha1.DatasetStatus{
				UfsTotal: "2Gi",
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "spark",
				Namespace: "fluid",
			},
			Status: datav1alpha1.DatasetStatus{
				UfsTotal: "",
			},
		},
	}

	testObjs := []runtime.Object{}
	for _, datasetInput := range datasetInputs {
		testObjs = append(testObjs, datasetInput.DeepCopy())
	}
	client := fake.NewFakeClientWithScheme(testScheme, testObjs...)

	engines := []JuiceFSEngine{
		{
			name:      "hbase",
			namespace: "fluid",
			Client:    client,
			Log:       log.NullLogger{},
		},
		{
			name:      "spark",
			namespace: "fluid",
			Client:    client,
			Log:       log.NullLogger{},
		},
	}

	for _, engine := range engines {
		err := engine.SyncMetadata()
		if err != nil {
			t.Errorf("fail to exec the function")
		}
	}

	engine := JuiceFSEngine{
		name:      "hadoop",
		namespace: "fluid",
		Client:    client,
		Log:       log.NullLogger{},
	}

	err := gohook.Hook(operations.JuiceFileUtils.QueryMetaDataInfoIntoFile, QueryMetaDataInfoIntoFileCommon, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = engine.SyncMetadata()
	if err != nil {
		t.Errorf("fail to exec function RestoreMetadataInternal")
	}
	wrappedUnhookQueryMetaDataInfoIntoFile()
}
