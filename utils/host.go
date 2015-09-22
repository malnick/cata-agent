package utils

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

func Memory() *mem.VirtualMemoryStat {
	memory, _ := mem.VirtualMemory()
	return memory
}

func Cpu() []cpu.CPUInfoStat {
	c, _ := cpu.CPUInfo()
	return c
}

func Load() *load.LoadAvgStat {
	l, _ := load.LoadAvg()
	return l
}
