// Copyright 2017 Xiaomi, Inc.
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

package funcs

import (
	"github.com/Taki-Kun/falcon-plus/common/model"
	"github.com/Taki-Kun/falcon-plus/modules/agent/g"
	"github.com/toolkits/nux"
	"github.com/Taki-Kun/process"
	"sync"
)

var (
	// processStatMapHistory = make(map[string][historyCount][]*process.ProcStat)
	processStatMapHistory   = make([]map[string][]*process.ProcStat, historyCount)
	processLock             = new(sync.RWMutex)
)

func UpdateProcessStats() (err error) {
	var pidsMap map[string][]int32
	pidsMap, err = findPid()
	if err != nil {
		return
	}

	processLock.Lock()
	defer processLock.Unlock()

	var processesMap = make(map[string][]*process.ProcStat)
	for tags, pids := range pidsMap {
		var processesStat []*process.ProcStat
		for _, pid := range pids {
			var processStat *process.ProcStat
			processStat, err = process.ProcessInfo(pid)
			if err != nil {
				return
			}
			processesStat = append(processesStat, processStat)
		}
		processesMap[tags] = processesStat

	}
	for i := historyCount - 1; i > 0; i-- {
		processStatMapHistory[i] = processStatMapHistory[i-1]
	}
	processStatMapHistory[0] = processesMap
	return nil
}

func findPid() (pidsMap map[string][]int32, err error) {
	pidsMap = make(map[string][]int32)

	reportProcesses := g.ReportProcesses()
	sz := len(reportProcesses)
	if sz == 0 {
		return
	}

	var ps []*nux.Proc
	ps, err = nux.AllProcs()
	if err != nil {
		return
	}

	pslen := len(ps)

	for tags, m := range reportProcesses {
		for i := 0; i < pslen; i++ {
			if is_a(ps[i], m) {
				pidsMap[tags] = append(pidsMap[tags], int32(ps[i].Pid))
			}
		}
	}
	return
}


func ProcessPrepared(tags string) bool {
	processLock.RLock()
	defer processLock.RUnlock()
	return processStatMapHistory[0][tags] != nil
}


func ProcessMetrics() (L []*model.MetricValue) {

	for tags, processesStat := range processStatMapHistory[0] {
		if !ProcessPrepared(tags) {
			continue
		}
		var processMemRss uint64
		var processMemVms uint64
		var processMemSwap uint64
		var processMemData uint64
		var processMemStack uint64
		var processMemLocked uint64
		for _, processStat := range processesStat {
			processMemRss = processMemRss + processStat.MemInfo.RSS
			processMemVms = processMemVms + processStat.MemInfo.VMS
			processMemSwap = processMemSwap + processStat.MemInfo.Swap
			processMemData = processMemData + processStat.MemInfo.Data
			processMemStack = processMemStack + processStat.MemInfo.Stack
			processMemLocked = processMemLocked + processStat.MemInfo.Locked
		}
		L = append(L, GaugeValue(g.PROCESS_MEM_RSS, processMemRss, tags))
		L = append(L, GaugeValue(g.PROCESS_MEM_VMS, processMemVms, tags))
		L = append(L, GaugeValue(g.PROCESS_MEM_SWAP, processMemSwap, tags))
		L = append(L, GaugeValue(g.PROCESS_MEM_DATA, processMemData, tags))
		L = append(L, GaugeValue(g.PROCESS_MEM_STACK, processMemStack, tags))
		L = append(L, GaugeValue(g.PROCESS_MEM_LOCKED, processMemLocked, tags))
	}

	return
}