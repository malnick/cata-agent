package main

import (
	"encoding/json"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	host "github.com/malnick/cata-agent/host"
	"net/http"
	//	"gopkg.in/yaml.v2"
)

// Define flags
var verbose = flag.Bool("v", false, "Define verbose logging output.")

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
	p.Hostname = host.Hostname()
	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Debug("failed to encode json for hostname")
		panic(err)
	}
}

// The root index
func Index(w http.ResponseWriter, r *http.Request) {
	log.Debug("Request for / index")
	fmt.Fprintf(w, "Cata Agent Started on /")
}

func main() {
	flag.Parse()
	// Debug level
	log.SetLevel(log.DebugLevel)
	if *verbose {
		log.Debug("Loglevel: Debug")
	} else {
		log.SetLevel(log.InfoLevel)
		log.Info("Loglevel: Info")
	}

	// Define our router object
	router := mux.NewRouter().StrictSlash(true)
	// Root index
	router.HandleFunc("/", Index)
	// Host index
	router.HandleFunc("/host", indexHost)
	// /container index
	router.HandleFunc("/host/hostname", routeHostname)
	// Handle a failover
	log.Fatal(http.ListenAndServe(":8080", router))
}
