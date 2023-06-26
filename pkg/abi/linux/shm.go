// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linux

import "math"

// shmat(2) flags. Source: include/uapi/linux/shm.h
const (
	SHM_RDONLY = 0o10000  // Read-only access.
	SHM_RND    = 0o20000  // Round attach address to SHMLBA boundary.
	SHM_REMAP  = 0o40000  // Take-over region on attach.
	SHM_EXEC   = 0o100000 // Execution access.
)

// IPCPerm.Mode upper byte flags. Source: include/linux/shm.h
const (
	SHM_DEST      = 0o1000  // Segment will be destroyed on last detach.
	SHM_LOCKED    = 0o2000  // Segment will not be swapped.
	SHM_HUGETLB   = 0o4000  // Segment will use huge TLB pages.
	SHM_NORESERVE = 0o10000 // Don't check for reservations.
)

// Additional Linux-only flags for shmctl(2). Source: include/uapi/linux/shm.h
const (
	SHM_LOCK   = 11
	SHM_UNLOCK = 12
	SHM_STAT   = 13
	SHM_INFO   = 14
)

// SHM defaults as specified by linux. Source: include/uapi/linux/shm.h
const (
	SHMMIN = 1
	SHMMNI = 4096
	SHMMAX = math.MaxUint64 - 1<<24
	SHMALL = math.MaxUint64 - 1<<24
	SHMSEG = 4096
)

// ShmidDS is equivalent to struct shmid64_ds. Source:
// include/uapi/asm-generic/shmbuf.h
//
// +marshal
type ShmidDS struct {
	ShmPerm    IPCPerm
	ShmSegsz   uint64
	ShmAtime   TimeT
	ShmDtime   TimeT
	ShmCtime   TimeT
	ShmCpid    int32
	ShmLpid    int32
	ShmNattach uint64

	Unused4 uint64
	Unused5 uint64
}

// ShmParams is equivalent to struct shminfo. Source: include/uapi/linux/shm.h
//
// +marshal
type ShmParams struct {
	ShmMax uint64
	ShmMin uint64
	ShmMni uint64
	ShmSeg uint64
	ShmAll uint64
}

// ShmInfo is equivalent to struct shm_info. Source: include/uapi/linux/shm.h
//
// +marshal
type ShmInfo struct {
	UsedIDs       int32 // Number of currently existing segments.
	_             [4]byte
	ShmTot        uint64 // Total number of shared memory pages.
	ShmRss        uint64 // Number of resident shared memory pages.
	ShmSwp        uint64 // Number of swapped shared memory pages.
	SwapAttempts  uint64 // Unused since Linux 2.4.
	SwapSuccesses uint64 // Unused since Linux 2.4.
}
