package utils

import (
	log "github.com/Sirupsen/logrus"
	"github.com/shirou/gopsutil/mem"
	"io/ioutil"
)

type Mem struct {
	Free           float64 `json:"free"`
	Total          float64 `json:"total"`
	PercentageUsed float64 `json:"percentage_used"`
}

type CpuInfo struct {
	Total string
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
	//	freeMatch, _ := regexp.Compile("MemFree:*")
	//	totalMatch, _ := regexp.Compile("MemTotal:*")
	//	file, err := os.Open("/proc/meminfo")
	//	if err != nil {
	//		log.Debug("Failed to open /proc/meminfo")
	//		msg := "Failed to open /proc/meminfo"
	//		memory.Free = msg
	//		memory.Total = msg
	//		return memory
	//	}
	//	defer file.Close()
	//
	//	scanner := bufio.NewScanner(file)
	//	for scanner.Scan() {
	//		if freeMatch.MatchString(scanner.Text()) {
	//			memory.Free = strings.Trim(strings.Split(scanner.Text(), ":")[1], " ")
	//		}
	//		if totalMatch.MatchString(scanner.Text()) {
	//			memory.Total = strings.Trim(strings.Split(scanner.Text(), ":")[1], " ")
	//		}
	//	}
	//	return memory
	v, _ := mem.VirtualMemory()
	memory.Free = v.Free
	memory.Total = v.Total
	memory.PercentageUsed = v.UsedPercent
	log.Debug(memory)
	return memory
}

func Cpu() (cpu CpuInfo) {
	cpu.Total = "null"
	return
}
