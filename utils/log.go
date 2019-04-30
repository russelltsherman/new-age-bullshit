package utils

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// SetVerbosity -
func SetVerbosity(count string) {
	var logLevel logrus.Level

	switch count {
	case "1":
		logLevel = logrus.ErrorLevel
	case "2":
		logLevel = logrus.WarnLevel
	case "3":
		logLevel = logrus.InfoLevel
	case "4":
		logLevel = logrus.DebugLevel
	case "5":
		logLevel = logrus.TraceLevel
	default:
		logLevel = logrus.FatalLevel
	}

	log.SetLevel(logLevel)
	log.Infof("Log verbosity set to %v\n", logLevel)
}
