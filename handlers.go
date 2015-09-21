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
		panic(err)
	}
}

// Memory route
func hostMemory(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request to /host/memory route")
}

// The root index
func Index(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request for / index")
	fmt.Fprintf(w, "Cata Agent Started on /")
}
