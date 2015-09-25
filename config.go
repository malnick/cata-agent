package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"regexp"
	"strings"
)

type Config struct {
	Name        string   `json:"name"`
	LogLevel    string   `json:"log_level"`
	Consoles    []string `json:"consoles"`
	ConsolePort string   `json:"console_port"`
	SplayTime   string   `json:"splay_time"`
}

func ParseEnv(c Config) Config {
	// Create a few matches for our env parsing down the road
	matchConsole, _ := regexp.Compile("CATA_CONSOLES=*")
	matchPort, _ := regexp.Compile("CATA_CONSOLE_PORT=*")
	matchSplay, _ := regexp.Compile("CATA_SPLAY_TIME=*")
	// Parse the env for our config
	for _, e := range os.Environ() {

		// Get the consoles from the env
		if matchConsole.MatchString(e) {
			log.Debug("Found Cata Console Match: ", e)
			consolesAry := strings.Split(strings.Split(e, "=")[1], ",")
			for _, console := range consolesAry {
				c.Consoles = append(c.Consoles, string(console))
			}
		}
		// get ports from the env
		if matchPort.MatchString(e) {
			log.Debug("Port match found: ", e)
			c.ConsolePort = strings.Split(e, "=")[1]
		}
		if matchSplay.MatchString(e) {
			log.Debug("Splay time found in ENV: ", e)
			c.SplayTime = strings.Split(e, "=")[1]
		}
	}
	// If the consoles were not passed, set a default
	if len(c.Consoles) < 1 {
		log.Debug("Consoles not found, setting default to localhost")
		c.Consoles = append(c.Consoles, "localhost")
	}
	// if ports not passed, set a default
	if len(c.ConsolePort) < 1 {
		log.Debug("Port match empty, setting default to 9000")
		c.ConsolePort = "9000"
	}
	if len(c.SplayTime) > 0 {
		log.Debug("Splay not found in env, setting default to 2m")
		c.SplayTime = "2m"
	}
	return c
}

func ParseConfig() (c Config) {
	log.SetLevel(log.DebugLevel)
	if *verbose {
		log.Debug("Loglevel: Debug")
		c.LogLevel = "Verbose"
	} else if os.Getenv("VERBOSE") == "true" {
		log.Debug("LogLevel: Debug")
		c.LogLevel = "Verbose"
	} else {
		log.SetLevel(log.InfoLevel)
		log.Info("Loglevel: Info")
		c.LogLevel = "Info"
	}
	c = ParseEnv(c)
	return c
}
