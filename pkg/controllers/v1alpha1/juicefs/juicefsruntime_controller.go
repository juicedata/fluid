/*

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
	"github.com/fluid-cloudnative/fluid/pkg/controllers"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	datav1alpha1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
)

var _ controllers.RuntimeReconcilerInterface = (*JuiceFSRuntimeReconciler)(nil)

// JuiceFSRuntimeReconciler reconciles a JuiceFSRuntime object
type JuiceFSRuntimeReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	*controllers.RuntimeReconciler
}

// NewRuntimeReconciler create controller for watching runtime custom resources created
func NewRuntimeReconciler(client client.Client,
	log logr.Logger,
	scheme *runtime.Scheme,
	recorder record.EventRecorder) *JuiceFSRuntimeReconciler {
	r := &JuiceFSRuntimeReconciler{
		Scheme: scheme,
	}
	r.RuntimeReconciler = controllers.NewRuntimeReconciler(r, client, log, recorder)
	return r
}

//+kubebuilder:rbac:groups=data.fluid.io,resources=juicefsruntimes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=data.fluid.io,resources=juicefsruntimes/status,verbs=get;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the JuiceFSRuntime object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.6.4/pkg/reconcile
func (r *JuiceFSRuntimeReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("juicefsruntime", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *JuiceFSRuntimeReconciler) SetupWithManager(mgr ctrl.Manager, options controller.Options) error {
	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(options).
		For(&datav1alpha1.JuiceFSRuntime{}).
		Complete(r)
}
