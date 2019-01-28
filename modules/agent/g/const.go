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

package g

import (
	"time"
)

// changelog:
// 3.1.3: code refactor
// 3.1.4: bugfix ignore configuration
// 5.0.0: 支持通过配置控制是否开启/run接口；收集udp流量数据；du某个目录的大小
// 5.1.0: 同步插件的时候不再使用checksum机制
// 5.1.1: 修复往多个transfer发送数据的时候crash的问题
// 5.1.2: ignore mount point when blocks=0
const (
	VERSION          = "5.1.2"
	COLLECT_INTERVAL = time.Second
	URL_CHECK_HEALTH = "url.check.health"
	NET_PORT_LISTEN  = "net.port.listen"
	DU_BS            = "du.bs"
	PROC_NUM         = "proc.num"
	PROCESS_CPU_BUSY_PERCENT  = "process.cpu.busy.percent"
	PROCESS_MEM_RSS  = "process.mem.rss"
	PROCESS_MEM_VMS  = "process.mem.vms"
	PROCESS_MEM_SWAP  = "process.mem.swap"
	PROCESS_MEM_DATA  = "process.mem.data"
	PROCESS_MEM_STACK  = "process.mem.stack"
	PROCESS_MEM_LOCKED  = "process.mem.locked"
	PROCESS_THREADS_NUMBER  = "process.threads.number"
	PROCESS_FD_NUMBER  = "process.fd.number"
	PROCESS_CTXSWITCHESVOLUNTARY  = "process.ctxSwitches.voluntary"
	PROCESS_CTXSWITCHESINVOLUNTARY  = "process.ctxSwitches.involuntary"
	PROCESS_IOCOUNTERS_READCOUNT  = "process.iocounters.readcount"
	PROCESS_IOCOUNTERS_WRITECOUNT  = "process.iocounters.writecount"
	PROCESS_IOCOUNTERS_READBYTES  = "process.iocounters.readbytes"
	PROCESS_IOCOUNTERS_WRITEBYTES  = "process.iocounters.writebytes"
	PROCESS_NETIOCOUNTERS_BYTESSENT = "process.netiocounters.bytesSent"
	PROCESS_NETIOCOUNTERS_BYTESRECV = "process.netiocounters.bytesRecv"
	PROCESS_NETIOCOUNTERS_PACKETSSENT = "process.netiocounters.packetsSent"
	PROCESS_NETIOCOUNTERS_PACKETSRECV = "process.netiocounters.packetsRecv"
	PROCESS_NETIOCOUNTERS_ERRIN = "process.netiocounters.errin"
	PROCESS_NETIOCOUNTERS_ERROUT = "process.netiocounters.errout"
	PROCESS_NETIOCOUNTERS_DROPIN = "process.netiocounters.dropin"
	PROCESS_NETIOCOUNTERS_DROPOUT = "process.netiocounters.dropout"
	PROCESS_NETIOCOUNTERS_FIFOIN = "process.netiocounters.fifoin"
	PROCESS_NETIOCOUNTERS_FIFOOUT = "process.netiocounters.fifoout"
	PROCESS_NETIOCOUNTERS_NUMBER = "process.connections.number"
)
