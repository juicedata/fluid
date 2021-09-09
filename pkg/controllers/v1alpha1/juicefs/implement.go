package juicefs

import (
	"github.com/fluid-cloudnative/fluid/pkg/ddc/base"
	cruntime "github.com/fluid-cloudnative/fluid/pkg/runtime"
)

func (r *JuiceFSRuntimeReconciler) GetOrCreateEngine(ctx cruntime.ReconcileRequestContext) (engine base.Engine, err error) {
	panic("implement me")
}

func (r *JuiceFSRuntimeReconciler) RemoveEngine(ctx cruntime.ReconcileRequestContext) {
	panic("implement me")
}
