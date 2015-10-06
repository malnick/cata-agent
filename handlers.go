package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	//	utils "github.com/malnick/cata-agent/utils"
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
	hostinfo := HostInfo()
	if err := json.NewEncoder(w).Encode(hostinfo); err != nil {
		log.Debug("Failed to encode host info to json")
	}
}

// Memory route
func routeHostMemory(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request to /host/memory route")
	mem := HostMemory()
	if err := json.NewEncoder(w).Encode(mem); err != nil {
		log.Debug("Failed to encode json for memory")
	}
}

// CPU route
func routeHostCpu(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request to /host/cpu route")
	cpu := HostCpu()
	if err := json.NewEncoder(w).Encode(cpu); err != nil {
		log.Debug("Failed to encode json for cpu")
	}
}

// Host load
func routeHostLoad(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request to /host/load route")
	load := HostLoad()
	if err := json.NewEncoder(w).Encode(load); err != nil {
		log.Debug("Failed to encode json for load")
	}
}

// Host disk
func routeHostDisk(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request to /host/disk route")
	disk := HostDisk()
	if err := json.NewEncoder(w).Encode(disk); err != nil {
		log.Debug("Failed to encode json for load")
	}
}

// Host netio
func routeHostNetio(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request to /host/netio route")
	net := HostNetio()
	if err := json.NewEncoder(w).Encode(net); err != nil {
		log.Debug("Failed to encode json for load")
	}
}

// Host netcon
func routeHostNetcon(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request to /host/netio route")
	net := HostNetcon()
	if err := json.NewEncoder(w).Encode(net); err != nil {
		log.Debug("Failed to encode json for load")
	}
}

// The root index
func Index(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request for / index")
	config := ParseConfig()
	if err := json.NewEncoder(w).Encode(config); err != nil {
		log.Debug("Failed to encode configuration for kata agent")
	}
}
