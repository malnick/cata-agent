package main

import (
	log "github.com/Sirupsen/logrus"
)

type Config struct {
	Name     string `json:"name"`
	LogLevel string `json:"log_level"`
}

func ParseConfig() (c Config) {
	log.SetLevel(log.DebugLevel)
	if *verbose {
		log.Debug("Loglevel: Debug")
		c.LogLevel = "Verbose"

	} else {
		log.SetLevel(log.InfoLevel)
		log.Info("Loglevel: Info")
		c.LogLevel = "Info"

	}
	c.Name = "Cata Agent : Configuration"
	return c
}
