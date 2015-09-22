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
	Name     string   `json:"name"`
	LogLevel string   `json:"log_level"`
	Alarms   []Alarm  `json:"alarms"`
	Consoles []string `json:"consoles"`
}

func GetAlarms() (alarms []Alarm) {
	matchEnv, _ := regexp.Compile("CATA_ALARM_*")
	for _, e := range os.Environ() {
		if matchEnv.MatchString(e) {
			var newAlarm Alarm
			log.Debug("New alarm: ", e)
			newAlarm.Name = strings.Split(strings.Split(e, "=")[0], "CATA_ALARM_")[1]
			// Crit value is first in list
			crit, _ := strconv.Atoi(strings.Split(strings.Split(e, "=")[1], ",")[0])
			warn, _ := strconv.Atoi(strings.Split(strings.Split(e, "=")[1], ",")[1])
			ok, _ := strconv.Atoi(strings.Split(strings.Split(e, "=")[1], ",")[2])
			newAlarm.Critical = crit
			newAlarm.Warning = warn
			newAlarm.Ok = ok
			alarms = append(alarms, newAlarm)
		}
	}
	return alarms
}

func GetConsoles() (consoles []string) {
	matchConsole, _ := regexp.Compile("CATA_CONSOLES=*")
	for _, c := range os.Environ() {
		if matchConsole.MatchString(c) {
			consolesAry := strings.Split(strings.Split(c, "=")[1], ",")
			for _, console := range consolesAry {
				consoles = append(consoles, string(console))
			}
		}
	}
	if len(consoles) < 1 {
		consoles = append(consoles, "localhost")
	}
	return
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
	c.Name = "Cata Agent : Configuration"

	c.Alarms = GetAlarms()

	c.Consoles = GetConsoles()

	return c
}
