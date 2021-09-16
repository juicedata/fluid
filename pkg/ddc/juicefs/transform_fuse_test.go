package juicefs

import (
	datav1alpha1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"testing"
)

func TestTransformFuse(t *testing.T) {
	juicefsSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "fluid",
		},
	}
	testObjs := []runtime.Object{}
	testObjs = append(testObjs, (*juicefsSecret).DeepCopy())

	client := fake.NewFakeClientWithScheme(testScheme, testObjs...)
	engine := JuiceFSEngine{
		name:      "test",
		namespace: "fluid",
		Client:    client,
		Log:       log.NullLogger{},
		runtime: &datav1alpha1.JuiceFSRuntime{
			Spec: datav1alpha1.JuiceFSRuntimeSpec{
				Fuse: datav1alpha1.JuiceFSFuseSpec{},
			},
		},
	}

	var tests = []struct {
		runtime      *datav1alpha1.JuiceFSRuntime
		dataset      *datav1alpha1.Dataset
		juicefsValue *JuiceFS
		expect       string
	}{
		{&datav1alpha1.JuiceFSRuntime{
			Spec: datav1alpha1.JuiceFSRuntimeSpec{
				Fuse: datav1alpha1.JuiceFSFuseSpec{
					SecretName: "test",
				},
			},
		}, &datav1alpha1.Dataset{
			Spec: datav1alpha1.DatasetSpec{
				Mounts: []datav1alpha1.Mount{{
					MountPoint: "local:///mnt/test",
					Name:       "test",
				}},
			}}, &JuiceFS{}, ""},
	}
	for _, test := range tests {
		err := engine.transformFuse(test.runtime, test.dataset, test.juicefsValue)
		if err != nil {
			t.Errorf("Got err %v", err)
		}
	}
}
