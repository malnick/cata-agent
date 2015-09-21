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
	flag.Parse()
	// Debug level
	log.SetLevel(log.DebugLevel)
	if *verbose {
		log.Debug("Loglevel: Debug")
	} else {
		log.SetLevel(log.InfoLevel)
		log.Info("Loglevel: Info")
	}

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
