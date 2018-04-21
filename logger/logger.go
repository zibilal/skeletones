package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"fmt"
	"io"
)

var isDebug bool

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	if os.Getenv("IS_DEBUG") == "" || os.Getenv("IS_DEBUG") == "1" {
		isDebug = true
	} else if os.Getenv("IS_DEBUG") == "0"{
		isDebug = false
	}
}

func getCallerInfo() (file string, line int, ok bool) {
	_, file, line, ok = runtime.Caller(2)

	return
}

func SetOutput(output io.Writer) {
	log.SetOutput(output)
}

func Debug(msg ...string) {
	if !isDebug {
		return
	}

	fname, ln, ok := getCallerInfo()
	if ok {
		log.WithFields(log.Fields{
			"on": fmt.Sprintf("%s:%d", fname, ln),
		}).Debug(msg)
	} else {
		log.Debug(msg)
	}
}

func Info(msg ...string) {
	if !isDebug {
		return
	}

	fname, ln, ok := getCallerInfo()
	if ok {
		log.WithFields(log.Fields{
			"on": fmt.Sprintf("%s:%d", fname, ln),
		}).Info(msg)
	} else {
		log.Info(msg)
	}
}

func Warn(msg ...string) {
	if !isDebug {
		return
	}

	fname, ln, ok := getCallerInfo()
	if ok {
		log.WithFields(log.Fields{
			"on": fmt.Sprintf("%s:%d", fname, ln),
		}).Warn(msg)
	} else {
		log.Warn(msg)
	}
}

func Fatal(msg ...string) {
	if !isDebug {
		return
	}

	fname, ln, ok := getCallerInfo()
	if ok {
		log.WithFields(log.Fields{
			"on": fmt.Sprintf("%s:%d", fname, ln),
		}).Fatal(msg)
	} else {
		log.Fatal(msg)
	}
}

func Error(msg ...string) {
	if !isDebug {
		return
	}

	fname, ln, ok := getCallerInfo()
	if ok {
		log.WithFields(log.Fields{
			"on": fmt.Sprintf("%s:%d", fname, ln),
		}).Error(msg)
	} else {
		log.Error(msg)
	}
}