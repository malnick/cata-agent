package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"net/http"
	//	"gopkg.in/yaml.v2"
)

// Define flags
var verbose = flag.Bool("v", false, "Define verbose logging output.")

func main() {
	// Parse flags
	flag.Parse()
	// Parse the config here before doing anything else
	_ = ParseConfig()
	// Run the router
	router := NewRouter()
	// Handle a failure
	log.Fatal(http.ListenAndServe(":8080", router))
}
