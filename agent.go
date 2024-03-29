package main

import (
	"bytes"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/malnick/gopsutil/cpu"
	"github.com/malnick/gopsutil/disk"
	"github.com/malnick/gopsutil/docker"
	"github.com/malnick/gopsutil/host"
	"github.com/malnick/gopsutil/load"
	"github.com/malnick/gopsutil/mem"
	"github.com/malnick/gopsutil/net"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type postData struct {
	Memory *mem.VirtualMemoryStat  `json:"memory"`
	CPU    []cpu.CPUInfoStat       `json:"cpu"`
	Host   *host.HostInfoStat      `json:"host"`
	Load   *load.LoadAvgStat       `json:"load"`
	Disk   *disk.DiskUsageStat     `json:"disk"`
	Netio  []net.NetIOCountersStat `json:"netio"`
	Netcon []net.NetConnectionStat `json:"netcon"`
	Docker map[string]string       `json:"docker"`
}

func getData(c Config) postData {
	var p postData
	for _, check := range c.Checks {
		switch check {
		case "memory":
			mem, _ := mem.VirtualMemory()
			p.Memory = mem
		case "cpu":
			cpu, _ := cpu.CPUInfo()
			p.CPU = cpu
		case "host":
			host, _ := host.HostInfo()
			p.Host = host
		case "load":
			load, _ := load.LoadAvg()
			p.Load = load
			//dockerids, _ := docker.GetDockerIDList()
		case "disk":
			disk, _ := disk.DiskUsage("/")
			p.Disk = disk
		case "netio":
			net, _ := net.NetIOCounters(true)
			p.Netio = net
		case "netcon":
			net, _ := net.NetConnections("inet")
			p.Netcon = net
		case "docker":
			returnDockerData := make(map[string]string)
			dock, _ := docker.GetDockerIDList()
			for _, id := range dock {
				v, err := docker.CgroupCPUDocker(id)
				if err != nil {
					log.Error("error %v", err)
				}
				if v.CPU == "" {
					log.Error("could not get CgroupCPU %v", v)
				}
				returnDockerData[id] = v.CPU
			}
			p.Docker = returnDockerData
		}
	}
	return p
}

func startAgent(c Config) {
	for {
	Inner:
		for _, console := range c.Consoles {
			url := strings.Join([]string{"http://", console, ":", c.ConsolePort, "/agent"}, "")
			log.Debug("POSTing to URL: ", url)
			// JSON Post
			json, _ := json.Marshal(getData(c))
			log.Debug("POST ", string(json))
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Warn("Failed to POST to ", url)
				continue Inner
			}
			defer resp.Body.Close() // Sleep for 1 minute before next POST
			log.Info("Successful POST to ", url)
			log.Debug("Response Status: ", resp.Status)
			log.Debug("Response Headers: ", resp.Header)
			body, _ := ioutil.ReadAll(resp.Body)
			log.Debug("Response Body: ", body)
		}
		time.Sleep(2 * time.Minute)
	}
}
