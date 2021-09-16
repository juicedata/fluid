package juicefs

import (
	"github.com/fluid-cloudnative/fluid/pkg/common/deprecated"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
)

func (j *JuiceFSEngine) getDeprecatedCommonLabelName() string {
	return deprecated.LabelAnnotationStorageCapacityPrefix + j.namespace + "-" + j.name
}

func (j *JuiceFSEngine) HasDeprecatedCommonLabelName() (deprecated bool, err error) {
	// return deprecated.LabelAnnotationStorageCapacityPrefix + e.namespace + "-" + e.name

	var (
		fuseName string = j.getFuseDaemonsetName()
		namespace  string = j.namespace
	)

	fuses, err := j.getDaemonset(fuseName, namespace)
	if err != nil {
		if apierrs.IsNotFound(err) {
			j.Log.Info("Fuses with deprecated label not found")
			deprecated = false
			err = nil
			return
		}
		j.Log.Error(err, "Failed to get fuse", "fuseName", fuseName)
		return deprecated, err
	}

	nodeSelectors := fuses.Spec.Template.Spec.NodeSelector
	j.Log.Info("The current node selectors for worker", "fuseName", fuseName, "nodeSelector", nodeSelectors)

	if _, deprecated = nodeSelectors[j.getDeprecatedCommonLabelName()]; deprecated {
		j.Log.Info("the deprecated node selector exists", "nodeselector", j.getDeprecatedCommonLabelName())
	} else {
		j.Log.Info("The deprecated node selector doesn't exist", "nodeselector", j.getDeprecatedCommonLabelName())
	}

	return
}
