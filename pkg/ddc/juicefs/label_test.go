package juicefs

import (
	"testing"
)

func TestJuiceFSEngine_getCommonLabelName(t *testing.T) {
	testCases := []struct {
		name      string
		namespace string
		out       string
	}{
		{
			name:      "fuse1",
			namespace: "fluid",
			out:       "fluid.io/s-fluid-fuse1",
		},
		{
			name:      "fuse2",
			namespace: "fluid",
			out:       "fluid.io/s-fluid-fuse2",
		},
		{
			name:      "common",
			namespace: "default",
			out:       "fluid.io/s-default-common",
		},
	}
	for _, testCase := range testCases {
		engine := &JuiceFSEngine{
			name:      testCase.name,
			namespace: testCase.namespace,
		}
		out := engine.getCommonLabelName()
		if out != testCase.out {
			t.Errorf("in: %s-%s, expect: %s, got: %s", testCase.namespace, testCase.name, testCase.out, out)
		}
	}
}
