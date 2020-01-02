package main

import (
	"github.com/bogowroc/go-playground/sandpit/03_projects/logging/log"
	"time"
)

func main() {
	log.Debug(true)

	log.Log("This is a debug statement...")

	logger := log.New(time.RFC3339, true)

	logger.Log("debug", "This is a debug statement...")
	logger.Log("warning", "no tasks found")
	logger.Log("error", "exiting: no work performed")
}
