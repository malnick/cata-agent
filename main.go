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

func indexContainer(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting /container index")
	fmt.Println("Container Index")
}

func indexHost(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting /host route")
	fmt.Println("Host index")
}

func routeHostname(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting /host/hostname route")
	var p struct {
		Hostname string `json:"hostname"`
	}
	p.Hostname = host.Hostname()
	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Debug("failed to encode json for hostname")
		panic(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting / index")
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
	log.Fatal(http.ListenAndServe(":8080", router))
	// Host index
	router.HandleFunc("/host", indexHost)
	// /container index
	http.HandleFunc("/host/hostname", routeHostname)
}
