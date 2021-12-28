package config

import (
	"os"
	"strconv"
	"time"
)

type config struct {
	Service string

	LogPath string

	PublishInterval time.Duration
	PublishAmount   int
	PublisherId     string
}

var Values = &config{}

func init() {

	// Microservice Id
	Values.Service = "events.redis.publisher"

	// Log file path
	Values.LogPath = os.Getenv("LOG_FILE_PATH")

	if Values.LogPath == "" {
		Values.LogPath = "./log/app.log"
	}

	PublishInterval := os.Getenv("PUBLISH_INTERVAL_MLSC")
	Values.PublishInterval = 1

	if PublishInterval != "" {
		intPublishInterval, _ := strconv.Atoi(PublishInterval)
		Values.PublishInterval = time.Duration(int64(intPublishInterval))
	}

	PublishAmount := os.Getenv("PUBLISH_AMOUNT")
	Values.PublishAmount = 0

	if PublishAmount != "" {
		intPublishAmount, _ := strconv.Atoi(PublishAmount)
		Values.PublishAmount = intPublishAmount
	}

	Values.PublisherId = os.Getenv("PUBLISHER_ID")

	if Values.PublisherId == "" {
		Values.PublisherId = "Golang-Publisher"
	}
}
