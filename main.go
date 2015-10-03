// +build linux darwin amd64
package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

// Define flags
var verbose = flag.Bool("v", false, "Define verbose logging output.")
var agentPort = flag.String("p", "8080", "Agent port")

func main() {
	// Parse flags
	flag.Parse()
	// Parse the config here before doing anything else
	config := ParseConfig()
	log.Info("Starting Cata Agent")
	// Start the agent
	go startAgent(config)
	// Run the router
	router := NewRouter()
	// Handle a failure
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *agentPort), router))
}
