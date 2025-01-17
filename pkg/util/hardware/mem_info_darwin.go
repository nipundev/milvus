// Copyright (C) 2019-2020 Zilliz. All rights reserved.
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

//go:build darwin || openbsd || freebsd
// +build darwin openbsd freebsd

package hardware

import (
	"github.com/shirou/gopsutil/v3/mem"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus/pkg/log"
)

// GetUsedMemoryCount returns the memory usage in bytes.
func GetUsedMemoryCount() uint64 {
	icOnce.Do(func() {
		ic, icErr = inContainer()
	})
	if icErr != nil {
		log.Error(icErr.Error())
		return 0
	}
	if ic {
		// in container, calculate by `cgroups`
		used, err := getContainerMemUsed()
		if err != nil {
			log.Warn("failed to get container memory used", zap.Error(err))
			return 0
		}
		return used
	}
	// not in container, calculate by `gopsutil`
	stats, err := mem.VirtualMemory()
	if err != nil {
		log.Warn("failed to get memory usage count",
			zap.Error(err))
		return 0
	}

	return stats.Used
}
