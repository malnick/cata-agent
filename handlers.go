package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	utils "github.com/malnick/cata-agent/utils"
	"net/http"
)

// The container index
func indexContainer(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request for /container index")
	fmt.Println("Container Index")
}

// The host index
func indexHost(w http.ResponseWriter, r *http.Request) {
	log.Debug("Requets for /host index")
	fmt.Println("Host index")
}

// The hostname route
func routeHostname(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request for /host/hostname route")
	// Hostname struct for return
	var p struct {
		Hostname string `json:"hostname"`
	}

	// Use our library to get the hostname of the current host
	p.Hostname = utils.Hostname()
	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Debug("failed to encode json for hostname")
	}
}

// Memory route
func routeHostMemory(w http.ResponseWriter, r *http.Request) {
	config := ParseConfig()
	log.Debug("Request to /host/memory route")
	mem := utils.Memory()
	for _, metric := range config.Alarms {
		if metric.Name == "MEMORY" {
			log.Debug("MEMORY Alarm found: ", metric, mem)
			mem.Alarm = MemoryAlarm(int(mem.PercentageUsed), metric)
		}
	}
	if err := json.NewEncoder(w).Encode(mem); err != nil {
		log.Debug("Failed to encode json for memory")
	}
}

// CPU route
func routeHostCpu(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request to /host/cpu route")
	cpu := utils.Cpu()
	if err := json.NewEncoder(w).Encode(cpu); err != nil {
		log.Debug("Failed to encode json for cpu")
	}
}

// Host load
func routeHostLoad(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request to /host/load route")
	load := utils.Load()
	if err := json.NewEncoder(w).Encode(load); err != nil {
		log.Debug("Failed to encode json for load")
	}
}

// The root index
func Index(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request for / index")
	config := ParseConfig()
	if err := json.NewEncoder(w).Encode(config); err != nil {
		log.Debug("Failed to encode configuration for cata agent")
	}
}
