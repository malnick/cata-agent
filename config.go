package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Alarm struct {
	Name     string `json:"name"`
	Critical int    `json:"critical"`
	Warning  int    `json:"warning"`
	Ok       int    `json:"ok"`
}

type Config struct {
	Name        string   `json:"name"`
	LogLevel    string   `json:"log_level"`
	Alarms      []Alarm  `json:"alarms"`
	Consoles    []string `json:"consoles"`
	ConsolePort string   `json:"console_port"`
}

func ParseEnv(c Config) Config {
	// Create a few matches for our env parsing down the road
	matchEnv, _ := regexp.Compile("CATA_ALARM_*")
	matchConsole, _ := regexp.Compile("CATA_CONSOLES=*")
	matchPort, _ := regexp.Compile("CATA_CONSOLE_PORT=*")
	// Parse the env for our config
	for _, e := range os.Environ() {
		if matchEnv.MatchString(e) {
			var newAlarm Alarm
			log.Debug("New alarm: ", e)
			newAlarm.Name = strings.Split(strings.Split(e, "=")[0], "CATA_ALARM_")[1]
			//	Crit value is first in list
			crit, _ := strconv.Atoi(strings.Split(strings.Split(e, "=")[1], ",")[0])
			warn, _ := strconv.Atoi(strings.Split(strings.Split(e, "=")[1], ",")[1])
			ok, _ := strconv.Atoi(strings.Split(strings.Split(e, "=")[1], ",")[2])
			newAlarm.Critical = crit
			newAlarm.Warning = warn
			newAlarm.Ok = ok
			c.Alarms = append(c.Alarms, newAlarm)
		}
		// Get the consoles from the env
		if matchConsole.MatchString(e) {
			consolesAry := strings.Split(strings.Split(e, "=")[1], ",")
			for _, console := range consolesAry {
				c.Consoles = append(c.Consoles, string(console))
			}
		}
		// If the consoles were not passed, set a default
		if len(c.Consoles) < 1 {
			c.Consoles = append(c.Consoles, "localhost")
		}
		// get ports from the env
		if matchPort.MatchString(e) {
			c.ConsolePort = strings.Split(e, "=")[1]
		}
		// if ports not passed, set a default
		if len(c.ConsolePort) < 1 {
			c.ConsolePort = "9000"
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
