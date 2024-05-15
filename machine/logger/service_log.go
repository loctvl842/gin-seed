package logger

import (
	"flag"
	"strings"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type Config struct {
	DefaultLevel string
	BasePrefix   string
	Env          string
}

type ServiceLogger interface {
	GetLogger(prefix string) Logger
}

type UTCFormatter struct {
	logrus.Formatter
}

// A default app logger
// Just write everything to console
type stdLogger struct {
	logger   *logrus.Logger
	cfg      Config
	logLevel string
	env      string
}

func NewAppLogService(config *Config) *stdLogger {
	if config == nil {
		config = &Config{}
	}

	if config.DefaultLevel == "" {
		config.DefaultLevel = "info"
	}

	logger := logrus.New()
	logger.Level = mustParseLevel(config.DefaultLevel)

	if config.Env == "" || config.Env == "dev" {
		logger.SetFormatter(&prefixed.TextFormatter{
			FullTimestamp: true,
		})
	} else {
		logger.SetFormatter(UTCFormatter{&logrus.JSONFormatter{}})
	}

	return &stdLogger{
		logger:   logger,
		cfg:      *config,
		logLevel: config.DefaultLevel,
	}
}

func (s *stdLogger) GetLogger(prefix string) Logger {
	var entry *logrus.Entry

	prefix = s.cfg.BasePrefix + "." + prefix
	prefix = strings.Trim(prefix, ".")

	if prefix == "" {
		entry = logrus.NewEntry(s.logger)
	} else {
		entry = s.logger.WithField("prefix", prefix)
	}

	l := &logger{entry}
	var log Logger = l

	return log
}

// Implement Runnable
func (s *stdLogger) Name() string { return "file-logger" }

func (s *stdLogger) InitFlags() {
	flag.StringVar(&s.logLevel, "log-level", s.logLevel, "Log level")
}

func (s *stdLogger) Configure() error {
	lv := mustParseLevel(s.logLevel)
	s.logger.SetLevel(lv)
	return nil
}

func (s *stdLogger) Run() error { return s.Configure() }

func (s *stdLogger) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}
