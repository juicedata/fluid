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

package v1alpha1

import (
	"github.com/fluid-cloudnative/fluid/pkg/common"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// JuiceFSRuntimeSpec defines the desired state of JuiceFSRuntime
type JuiceFSRuntimeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The version information that instructs fluid to orchestrate a particular version of JuiceFS.
	JuiceFSVersion VersionSpec `json:"juicefsVersion,omitempty"`

	// The spec of init users
	InitUsers InitUsersSpec `json:"initUsers,omitempty"`

	// Desired state for JuiceFS Fuse
	Fuse JuiceFSFuseSpec `json:"fuse,omitempty"`

	// Tiered storage used by JuiceFS
	TieredStore TieredStore `json:"tieredstore,omitempty"`

	// The replicas of the cache client, need to be specified
	Replicas int32 `json:"replicas,omitempty"`

	// Manage the user to run Juicefs Runtime
	RunAs *User `json:"runAs,omitempty"`

	// Disable monitoring for JuiceFS Runtime
	// Prometheus is enabled by default
	// +optional
	DisablePrometheus bool `json:"disablePrometheus,omitempty"`
}

type JuiceFSFuseSpec struct {
	// Image for JuiceFS fuse
	Image string `json:"image,omitempty"`

	// Image for JuiceFS fuse
	ImageTag string `json:"image_tag,omitempty"`

	// One of the three policies: `Always`, `IfNotPresent`, `Never`
	ImagePullPolicy string `json:"image_pull_policy,omitempty"`

	// Secret name which is used by JuiceFS fuse
	SecretName string `json:"secret_name"`

	// Environment variables that will be used by JuiceFS Fuse
	Env map[string]string `json:"env"`

	// Resources that will be requested by JuiceFS Fuse.
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// If the fuse client should be deployed in global mode,
	// otherwise the affinity should be considered
	// +optional
	Global bool `json:"global,omitempty"`

	// NodeSelector is a selector which must be true for the fuse client to fit on a node,
	// this option only effect when global is enabled
	// +optional
	NodeSelector map[string]string `json:"node_selector,omitempty"`
}

// JuiceFSRuntimeStatus defines the observed state of JuiceFSRuntime
type JuiceFSRuntimeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// config map used to set configurations

	// FusePhase is the Fuse running phase
	FusePhase RuntimePhase `json:"fusePhase"`

	// Reason for the condition's last transition.
	FuseReason string `json:"fuseReason,omitempty"`

	// The total number of nodes that can be running the runtime Fuse
	// pod (including nodes correctly running the runtime Fuse pod).
	CurrentFuseNumberScheduled int32 `json:"currentFuseNumberScheduled"`

	// The total number of nodes that should be running the runtime Fuse
	// pod (including nodes correctly running the runtime Fuse pod).
	DesiredFuseNumberScheduled int32 `json:"desiredFuseNumberScheduled"`

	// The number of nodes that should be running the runtime Fuse pod and have one
	// or more of the runtime Fuse pod running and ready.
	FuseNumberReady int32 `json:"fuseNumberReady"`

	// The number of nodes that should be running the
	// runtime fuse pod and have none of the runtime fuse pod running and available
	// (ready for at least spec.minReadySeconds)
	// +optional
	FuseNumberUnavailable int32 `json:"fuseNumberUnavailable,omitempty"`

	// The number of nodes that should be running the
	// runtime Fuse pod and have one or more of the runtime Fuse pod running and
	// available (ready for at least spec.minReadySeconds)
	// +optional
	FuseNumberAvailable int32 `json:"fuseNumberAvailable,omitempty"`

	// Duration tell user how much time was spent to setup the runtime
	SetupDuration string `json:"setupDuration,omitempty"`

	// Represents the latest available observations of a ddc runtime's current state.
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []RuntimeCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// CacheStatus represents the total resources of the dataset.
	CacheStates common.CacheStateList `json:"cacheStates,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// JuiceFSRuntime is the Schema for the juicefsruntimes API
type JuiceFSRuntime struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JuiceFSRuntimeSpec   `json:"spec,omitempty"`
	Status JuiceFSRuntimeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// JuiceFSRuntimeList contains a list of JuiceFSRuntime
type JuiceFSRuntimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JuiceFSRuntime `json:"items"`
}

func init() {
	SchemeBuilder.Register(&JuiceFSRuntime{}, &JuiceFSRuntimeList{})
}

// Replicas gets the replicas of runtime worker
func (r *JuiceFSRuntime) Replicas() int32 {
	return r.Spec.Replicas
}
