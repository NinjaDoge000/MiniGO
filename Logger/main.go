package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)

	log.Info("Started")
	log.WithField("url", "https://example.com").Info("Crawling")
	log.Warn("Missing H1 tag")
	log.Error("Request timeout")
}
