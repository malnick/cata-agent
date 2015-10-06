package main

import (
	//	log "github.com/Sirupsen/logrus"
	"github.com/malnick/gopsutil/cpu"
	"github.com/malnick/gopsutil/disk"
	"github.com/malnick/gopsutil/docker"
	"github.com/malnick/gopsutil/host"
	"github.com/malnick/gopsutil/load"
	"github.com/malnick/gopsutil/mem"
	"github.com/malnick/gopsutil/net"
)

func HostInfo() *host.HostInfoStat {
	h, _ := host.HostInfo()
	return h
}

func HostMemory() *mem.VirtualMemoryStat {
	memory, _ := mem.VirtualMemory()
	return memory
}

func HostCpu() []cpu.CPUInfoStat {
	c, _ := cpu.CPUInfo()
	return c
}

func HostLoad() *load.LoadAvgStat {
	l, _ := load.LoadAvg()
	return l
}

func HostDisk() *disk.DiskUsageStat {
	d, _ := disk.DiskUsage("/")
	return d
}

func HostNetio() []net.NetIOCountersStat {
	net, _ := net.NetIOCounters(true)
	return net
}

func HostNetcon() []net.NetConnectionStat {
	net, _ := net.NetConnections("inet")
	return net
}

func HostDocker() map[string]*cpu.CPUTimesStat {
	returnDockerData := make(map[string]*cpu.CPUTimesStat)
	dock, _ := docker.GetDockerIDList()
	for _, id := range dock {
		v, _ := docker.CgroupCPUDocker(id)
		returnDockerData[id] = v
	}
	return returnDockerData
}
