package utils

import (
	log "github.com/Sirupsen/logrus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"

	"io/ioutil"
)

type Mem struct {
	Free           uint64  `json:"free"`
	Total          uint64  `json:"total"`
	PercentageUsed float64 `json:"percentage_used"`
}

func Hosts() string {
	h, _ := ioutil.ReadFile("/etc/hosts")
	return string(h)
}

func Hostname() string {
	h, _ := ioutil.ReadFile("/etc/hostname")
	return string(h)
}

func Memory() (memory Mem) {
	v, _ := mem.VirtualMemory()
	memory.Free = v.Free
	memory.Total = v.Total
	memory.PercentageUsed = v.UsedPercent
	log.Debug(memory)
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