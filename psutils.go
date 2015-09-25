package main

import (
	//	log "github.com/Sirupsen/logrus"
	"github.com/malnick/gopsutil/cpu"
	"github.com/malnick/gopsutil/host"
	"github.com/malnick/gopsutil/load"
	"github.com/malnick/gopsutil/mem"
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
