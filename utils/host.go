package utils

import (
	log "github.com/Sirupsen/logrus"
	"github.com/malnick/gopsutil/cpu"
	"github.com/malnick/gopsutil/load"
	"github.com/malnick/gopsutil/mem"

	"io/ioutil"
)

type Mem struct {
	Free           uint64  `json:"free"`
	Total          uint64  `json:"total"`
	PercentageUsed float64 `json:"percentage_used"`
	Alarm          string  `json:"alarm"`
}

func Hosts() string {
	h, _ := ioutil.ReadFile("/etc/hosts")
	return string(h)
}

func Hostname() string {
	h, _ := ioutil.ReadFile("/etc/hostname")
	return string(h)
}

func Memory(alarm string) (memory Mem) {
	v, _ := mem.VirtualMemory()
	memory.Free = v.Free
	memory.Total = v.Total
	memory.PercentageUsed = v.UsedPercent
	if len(alarm) > 0 {
		memory.Alarm = MemoryAlarm(int(memory.PercentageUsed), alarm)
	}
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
