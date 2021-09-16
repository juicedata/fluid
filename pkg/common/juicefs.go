package common

// Runtime for Juice
const (
	JUICEFS_RUNTIME = "juicefs"

	JUICEFS_MOUNT_TYPE = "JuiceFS"

	JUICEFS_NAMESPACE = "juicefs-system"

	JUICEFS_CHART = JUICEFS_MOUNT_TYPE

	JUICEFS_RUNTIME_IMAGE_ENV = "JUICEFS_RUNTIME_IMAGE_ENV"

	JUICEFS_FUSE_IMAGE_ENV = "JUICEFS_FUSE_IMAGE_ENV"

	DEFAULT_JUICEFS_RUNTIME_IMAGE = "juicedata/juicefs-csi-driver:v0.10.5"

	DEFAULT_JUICEFS_FUSE_IMAGE = "juicedata/juicefs-csi-driver:v0.10.5"

	JuiceFSMountPath = "/bin/mount.juicefs"
)
