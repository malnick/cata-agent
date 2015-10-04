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
	EnableApi   string   `json:"enable_api"`
	Checks      []string `json:"checks"`
}

// Defaults
var DefaultConsoles = "localhost"
var DefaultPort = "9000"
var DefaultEnableApi = "true"
var DefaultSplayTime = "2m"
var DefaultChecks = []string{"memory", "host", "load", "disk", "netio", "netcon"}

// ParseEnv parses the agent environ and gets kata_* specific data
func ParseEnv(c Config) Config {
	// Set defaults, override later
	c.ConsolePort = DefaultPort
	c.EnableApi = DefaultEnableApi
	c.SplayTime = DefaultSplayTime
	c.Consoles = append(c.Consoles, "localhost")
	c.Checks = DefaultChecks
	// Name needs to be set to hostname - TODO
	//c.Name = getHostname()?

	// Create a few matches for our env parsing down the road
	matchConsole, _ := regexp.Compile("KATA_CONSOLES=*")
	matchPort, _ := regexp.Compile("KATA_CONSOLE_PORT=*")
	matchSplay, _ := regexp.Compile("KATA_SPLAY_TIME=*")
	matchApiBool, _ := regexp.Compile("KATA_ENABLE_API=*")
	matchChecks, _ := regexp.Compile("KATA_CHECKS=*")

	// Parse the env for our config
	log.Debug("Setting ENV configuration..")
	for _, e := range os.Environ() {

		// kata consoles to report to
		if matchConsole.MatchString(e) {
			log.Debug("Consoles: ", e)
			consolesAry := strings.Split(strings.Split(e, "=")[1], ",")
			for _, console := range consolesAry {
				c.Consoles = []string{string(console)}
			}
		}

		// kata console port
		if matchPort.MatchString(e) {
			log.Debug("Console Port: ", e)
			c.ConsolePort = strings.Split(e, "=")[1]
		}

		// Splay time configuration
		if matchSplay.MatchString(e) {
			log.Debug("Splay Time ENV: ", e)
			c.SplayTime = strings.Split(e, "=")[1]
		}

		// Enable agent REST API
		if matchApiBool.MatchString(e) {
			log.Debug("API Enable: ", e)
			c.EnableApi = strings.Split(e, "=")[1]
		}

		// Which checks to perform
		if matchChecks.MatchString(e) {
			log.Debug("Checks: ", e)
			for _, check := range strings.Split(strings.Split(e, "=")[1], ",") {
				log.Debug("Enabling check ", e)
				c.Checks = append(c.Checks, check)
			}
		}
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
