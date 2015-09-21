package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"regexp"
)

type Alarm struct {
	Name     string `json:"name"`
	Critical int    `json:"critical"`
	Warning  int    `json:"warning"`
	Ok       int    `json:"ok"`
	Unknown  int    `json:"unknown"`
}

type Config struct {
	Name     string  `json:"name"`
	LogLevel string  `json:"log_level"`
	Alarms   []Alarm `json:"alarms"`
}

func GetAlarms() (alarms []Alarm) {
	matchEnv, _ := regexp.Compile("CATA_ALARM_*")
	for _, e := range os.Environ() {
		if matchEnv.MatchString(e) {
			var newAlarm Alarm
			log.Debug("New alarm: ", e)
			newAlarm.Name = e
			// Add here
			alarms = append(alarms, newAlarm)
		}
	}
	return alarms
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

	return c
}
