/*
Copyright 2021 The Fluid Authors.

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

const (
	BlockCacheBytes     = "juicefs_blockcache_bytes"
	BlockCacheHits      = "juicefs_blockcache_hits"
	BlockCacheMiss      = "juicefs_blockcache_miss"
	BlockCacheHitBytes  = "juicefs_blockcache_hit_bytes"
	BlockCacheMissBytes = "juicefs_blockcache_miss_bytes"
	UsedSpace           = "juicefs_used_space"

	PodRoleType = "role"

	WorkerPodRole = "juicefs-worker"

	METADATA_SYNC_NOT_DONE_MSG                = "[Calculating]"
	CHECK_METADATA_SYNC_DONE_TIMEOUT_MILLISEC = 500
)