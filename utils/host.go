package utils

import (
	"io/ioutil"
	"regexp"
)

func Hosts() string {
	h, _ := ioutil.ReadFile("/etc/hosts")
	return string(h)
}

func Hostname() string {
	h, _ := ioutil.ReadFile("/etc/hostname")
	return string(h)
}

type Mem struct {
	Free  string
	Total string
}

func Memory() (memory Mem) {
	freeMatch := regexp.Compile("MemFree:*")
	totalMatch := regexp.Compile("MemTotal:*")
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Debug("Failed to open /proc/meminfo")
		return memory{
			"Failed to open /proc/meminfo",
			"Failed to open /proc/meminfo",
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if freeMatch.MatchString(scanner.Text()) {
			memory.Free = scanner.Text()
		}
		if totalMatch.MatchString(scanner.Text()) {
			memory.Total = scanner.Text()
		}
	}
	return memory
}
