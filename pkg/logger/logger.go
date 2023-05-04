package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var l = initLogger()

func initLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
}

func getLogger() *logrus.Entry {
	return l.WithField("time", fmt.Sprintf("[%s]", time.Now().UTC().String()))
}

func Infof(message string, args ...interface{}) {
	getLogger().Infof(message, args...)
}

func Errorf(message string, args ...interface{}) {
	getLogger().Errorf(message, args...)
}

func Debugf(message string, args ...interface{}) {
	getLogger().Debugf(message, args...)
}

func Fatalf(message string, args ...interface{}) {
	getLogger().Fatalf(message, args...)
}

func Info(message string) {
	getLogger().Info(message)
}
