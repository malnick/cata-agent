package main

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func startAgent(c Config) {
	log.Debug("Starting cata agent")
	for {
		for _, console := range c.Consoles {
			url := strings.Join([]string{"http://", console, ":", c.ConsolePort}, "")
			log.Debug("POSTing to URL: ", url)
			jsonStr := []byte(`{"title":"Buy cheese and bread for breakfast."}`)
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Warn("Failed to POST to ", url)
				break
			}
			defer resp.Body.Close() // Sleep for 1 minute before next POST
			log.Debug("Response Status: ", resp.Status)
			log.Debug("Response Headers: ", resp.Header)
			body, _ := ioutil.ReadAll(resp.Body)
			log.Debug("Response Body: ", body)
		}
		time.Sleep(time.Minute * 1)
	}
}
