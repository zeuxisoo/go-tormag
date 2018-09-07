package logger

import (
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

func init() {
	log.Formatter = new(logrus.TextFormatter)
	log.Level = logrus.DebugLevel
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}
