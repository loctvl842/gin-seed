package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type Fields logrus.Fields

var currentServLog ServiceLogger

func InitServLogger() {
	currentServLog = NewAppLogService(&Config{
		BasePrefix:   "core",
		DefaultLevel: "info",
		Env:          os.Getenv("APP_ENV"),
	})
}

func GetCurrent() ServiceLogger {
	return currentServLog
}

type Logger interface {
	Print(args ...interface{})
	Debug(...interface{})
	Debugln(...interface{})
	Debugf(string, ...interface{})

	Info(...interface{})
	Infoln(...interface{})
	Infof(string, ...interface{})

	Warn(...interface{})
	Warnln(...interface{})
	Warnf(string, ...interface{})

	Error(...interface{})
	Errorln(...interface{})
	Errorf(string, ...interface{})

	Fatal(...interface{})
	Fatalln(...interface{})
	Fatalf(string, ...interface{})

	Panic(...interface{})
	Panicln(...interface{})
	Panicf(string, ...interface{})

	With(key string, value interface{}) Logger
	Withs(Fields) Logger
	// add source field to log
	WithSrc() Logger
	GetLevel() string
}

type logger struct {
	*logrus.Entry
}

func (l *logger) GetLevel() string {
	return l.Logger.Level.String()
}

func (l *logger) debugSrc() *logrus.Entry {
	if _, ok := l.Data["source"]; ok {
		return l.Entry
	}

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		file = file[slash+1:]
	}
	return l.WithField("source", fmt.Sprintf("%s:%d", file, line))
}

func (l *logger) With(key string, value interface{}) Logger {
	return &logger{l.WithField(key, value)}
}

func (l *logger) Withs(fields Fields) Logger {
	return &logger{l.WithFields(logrus.Fields(fields))}
}

func (l *logger) WithSrc() Logger {
	return &logger{l.debugSrc()}
}

func mustParseLevel(level string) logrus.Level {
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		log.Fatal(err.Error())
	}
	return lv
}
