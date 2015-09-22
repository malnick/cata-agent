package utils

import (
	//	log "github.com/Sirupsen/logrus"
	"github.com/malnick/gopsutil/cpu"
	"github.com/malnick/gopsutil/host"
	"github.com/malnick/gopsutil/load"
	"github.com/malnick/gopsutil/mem"
)

type MemData struct {
	Free           uint64  `json:"free"`
	Total          uint64  `json:"total"`
	PercentageUsed float64 `json:"percentage_used"`
	Alarm          string  `json:"alarm"`
}

type CpuData struct {
	CpuInfo []cpu.CPUInfoStat `json:"cpu_info"`
}

type HostData struct {
	Data *host.HostInfoStat
}

func HostInfo() *host.HostInfoStat {
	h, _ := host.HostInfo()
	return h
}

func Memory() (memory MemData) {
	v, _ := mem.VirtualMemory()
	memory.Free = v.Free
	memory.Total = v.Total
	memory.PercentageUsed = v.UsedPercent
	return memory
}

func Cpu() (cpuData CpuData) {
	c, _ := cpu.CPUInfo()
	cpuData.CpuInfo = c
	return cpuData
}

func Load() *load.LoadAvgStat {
	l, _ := load.LoadAvg()
	return l
}
