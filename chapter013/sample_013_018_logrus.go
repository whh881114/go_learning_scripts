package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"tool": "pen",
	}).Info("This is a pen.")
}
