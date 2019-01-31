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

package cron

import (
	"github.com/Taki-Kun/falcon-plus/common/model"
	"github.com/Taki-Kun/falcon-plus/modules/agent/g"
	"log"
	"strconv"
	"strings"
	"time"
)

func SyncBuiltinMetrics() {
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		go syncBuiltinMetrics()
	}
}

func syncBuiltinMetrics() {

	var timestamp int64 = -1
	var checksum string = "nil"

	duration := time.Duration(g.Config().Heartbeat.Interval) * time.Second

	for {
		time.Sleep(duration)

		var ports = []int64{}
		var paths = []string{}
		var processes = make(map[string]map[int]string)
		var procs = make(map[string]map[int]string)
		var urls = make(map[string]string)

		hostname, err := g.Hostname()
		if err != nil {
			continue
		}

		req := model.AgentHeartbeatRequest{
			Hostname: hostname,
			Checksum: checksum,
		}

		var resp model.BuiltinMetricResponse
		err = g.HbsClient.Call("Agent.BuiltinMetrics", req, &resp)
		if err != nil {
			log.Println("ERROR:", err)
			continue
		}

		if resp.Timestamp <= timestamp {
			continue
		}

		if resp.Checksum == checksum {
			continue
		}

		timestamp = resp.Timestamp
		checksum = resp.Checksum

		for _, metric := range resp.Metrics {

			if metric.Metric == g.URL_CHECK_HEALTH {
				arr := strings.Split(metric.Tags, ",")
				if len(arr) != 2 {
					continue
				}
				url := strings.Split(arr[0], "=")
				if len(url) != 2 {
					continue
				}
				stime := strings.Split(arr[1], "=")
				if len(stime) != 2 {
					continue
				}
				if _, err := strconv.ParseInt(stime[1], 10, 64); err == nil {
					urls[url[1]] = stime[1]
				} else {
					log.Println("metric ParseInt timeout failed:", err)
				}
			}

			if metric.Metric == g.NET_PORT_LISTEN {
				arr := strings.Split(metric.Tags, "=")
				if len(arr) != 2 {
					continue
				}

				if port, err := strconv.ParseInt(arr[1], 10, 64); err == nil {
					ports = append(ports, port)
				} else {
					log.Println("metrics ParseInt failed:", err)
				}

				continue
			}

			if metric.Metric == g.DU_BS {
				arr := strings.Split(metric.Tags, "=")
				if len(arr) != 2 {
					continue
				}

				paths = append(paths, strings.TrimSpace(arr[1]))
				continue
			}

			if metric.Metric == g.PROCESS_MEM_RSS ||
				metric.Metric == g.PROCESS_MEM_VMS ||
				metric.Metric == g.PROCESS_MEM_SWAP ||
				metric.Metric == g.PROCESS_MEM_DATA ||
				metric.Metric == g.PROCESS_MEM_STACK ||
				metric.Metric == g.PROCESS_MEM_LOCKED ||
				metric.Metric == g.PROCESS_CPU_BUSY_PERCENT ||
				metric.Metric == g.PROCESS_THREADS_NUMBER ||
				metric.Metric == g.PROCESS_FD_NUMBER ||
				metric.Metric == g.PROCESS_CTXSWITCHESVOLUNTARY ||
				metric.Metric == g.PROCESS_CTXSWITCHESINVOLUNTARY ||
				metric.Metric == g.PROCESS_IOCOUNTERS_READCOUNT ||
				metric.Metric == g.PROCESS_IOCOUNTERS_WRITECOUNT ||
				metric.Metric == g.PROCESS_IOCOUNTERS_READBYTES ||
				metric.Metric == g.PROCESS_IOCOUNTERS_WRITEBYTES ||
				metric.Metric == g.PROCESS_NETIOCOUNTERS_BYTESSENT ||
				metric.Metric == g.PROCESS_NETIOCOUNTERS_BYTESRECV ||
				metric.Metric == g.PROCESS_NETIOCOUNTERS_PACKETSSENT ||
				metric.Metric == g.PROCESS_NETIOCOUNTERS_PACKETSRECV ||
				metric.Metric == g.PROCESS_NETIOCOUNTERS_ERRIN ||
				metric.Metric == g.PROCESS_NETIOCOUNTERS_ERROUT ||
				metric.Metric == g.PROCESS_NETIOCOUNTERS_DROPIN ||
				metric.Metric == g.PROCESS_NETIOCOUNTERS_DROPOUT ||
				metric.Metric == g.PROCESS_NETIOCOUNTERS_FIFOIN ||
				metric.Metric == g.PROCESS_NETIOCOUNTERS_FIFOOUT ||
				metric.Metric == g.PROCESS_NETIOCOUNTERS_NUMBER {
				arr := strings.Split(metric.Tags, ",")

				tmpMap := make(map[int]string)

				for i := 0; i < len(arr); i++ {
					if strings.HasPrefix(arr[i], "name=") {
						tmpMap[1] = strings.TrimSpace(arr[i][5:])
					} else if strings.HasPrefix(arr[i], "cmdline=") {
						tmpMap[2] = strings.TrimSpace(arr[i][8:])
					}
				}

				processes[metric.Tags] = tmpMap
				continue
			}

			if metric.Metric == g.PROC_NUM {
				arr := strings.Split(metric.Tags, ",")

				tmpMap := make(map[int]string)

				for i := 0; i < len(arr); i++ {
					if strings.HasPrefix(arr[i], "name=") {
						tmpMap[1] = strings.TrimSpace(arr[i][5:])
					} else if strings.HasPrefix(arr[i], "cmdline=") {
						tmpMap[2] = strings.TrimSpace(arr[i][8:])
					}
				}

				procs[metric.Tags] = tmpMap
			}
		}

		g.SetReportUrls(urls)
		g.SetReportPorts(ports)
		g.SetReportProcs(procs)
		g.SetDuPaths(paths)
		g.SetReportProcesses(processes)

	}
}
