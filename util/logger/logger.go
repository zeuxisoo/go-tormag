package logger

import (
    "github.com/sirupsen/logrus"
)

var (
    log = logrus.New()
)

func init() {
    log.Formatter = new(logrus.TextFormatter)
    log.Level     = logrus.DebugLevel
}

func Infof(format string, args ...interface{}) {
    log.Infof(format, args...)
}
