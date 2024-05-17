package gosdk

import (
	"app/addons/logger"
	"app/constant"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
)

type application struct {
	logger       logger.Logger
	signalChan   chan os.Signal
	cmdLine      *AppFlagSet
	initServices map[string]PrefixRunnable
	name         string
	version      string
	env          string
	opts         []Option
	subServices  []Runnable
	isRegister   bool
}

func (s *application) Name() string { return s.name }

func (s *application) Version() string { return s.version }

func (s *application) Init() error {
	for _, sv := range s.initServices {
		if err := sv.Run(); err != nil {
			return err
		}
	}

	return nil
}

func (s *application) IsRegistered() bool {
	return s.isRegister
}

func (s *application) Start() error {
	signal.Notify(s.signalChan, os.Interrupt)

	c := make(chan error, 1)

	// Run all sub services
	for _, sv := range s.subServices {
		go func(r Runnable) { c <- r.Run() }(sv)
	}

	for {
		select {
		case err := <-c:
			if err != nil {
				return err
			}
		// Handle signals
		case sig := <-s.signalChan:
			switch sig {
			case syscall.SIGHUP:
				fmt.Println("Reload configuration")
				return nil
			case syscall.SIGTERM:
				fmt.Println("Received SIGTERM, exiting")
				s.Stop()
				return nil
			case syscall.SIGINT:
				fmt.Println("Received SIGINT, exiting")
				s.Stop()
				return nil
			default:
				fmt.Println("Received signal: ", sig)
				s.Stop()
				return nil
			}
		}
	}
}

func (s *application) Stop() {
	s.logger.Infoln("Stopping addons...")
	var wg sync.WaitGroup

	stop := func(r Runnable) {
		defer wg.Done()
		<-r.Stop()
	}
	wg.Add(len(s.subServices) + len(s.initServices))

	// Stop all sub services
	for _, sv := range s.subServices {
		go stop(sv)
	}

	// Stop all init services
	for _, sv := range s.initServices {
		go stop(sv)
	}

	wg.Wait()
	s.logger.Info("All addons are stopped")
}

func (s *application) OutEnv() {
	s.cmdLine.GetSampleEnvs()
}

func (s *application) Logger(prefix string) logger.Logger {
	return s.logger
}

func (s *application) parseFlags() {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}
	_, err := os.Stat(envFile)
	if err == nil {
		err := godotenv.Load(envFile)
		if err != nil {
			s.logger.Fatalf("Loading env(%s): %s", envFile, err.Error())
		}
	} else if envFile != ".env" {
		s.logger.Fatalf("Loading env(%s): %s", envFile, err.Error())
	}

	s.cmdLine.Parse([]string{})
}

func (s *application) Get(prefix string) (interface{}, bool) {
	is, ok := s.initServices[prefix]
	if !ok {
		return nil, ok
	}
	return is.Get(), ok
}

func (s *application) MustGet(prefix string) interface{} {
	is, ok := s.initServices[prefix]
	if !ok {
		panic("MustGet: no such service")
	}
	return is.Get()
}

func (s *application) Env() string { return s.env }

func New(opts ...Option) Application {
	sv := &application{
		opts:         opts,
		signalChan:   make(chan os.Signal, 1),
		subServices:  []Runnable{},
		initServices: map[string]PrefixRunnable{},
	}

	// init default logger
	logger.InitServLogger()
	sv.logger = logger.GetCurrent().GetLogger("service")

	for _, opt := range opts {
		opt(sv)
	}

	sv.initFlags()

	loggerRunnable := logger.GetCurrent().(Runnable)
	loggerRunnable.InitFlags()

	sv.cmdLine = newFlagSet(sv.name, flag.CommandLine)
	sv.parseFlags()

	_ = loggerRunnable.Configure()

	return sv
}

func (s *application) initFlags() {
	flag.StringVar(&s.env, "env", string(constant.Dev), "Env for service. Ex: dev | stg | prd")

	for _, subService := range s.subServices {
		subService.InitFlags()
	}

	for _, initService := range s.initServices {
		initService.InitFlags()
	}
}

func WithName(name string) Option {
	return func(s *application) { s.name = name }
}

func WithVersion(version string) Option {
	return func(s *application) { s.version = version }
}

// Add init component to SDK
// These components will run sequentially before service run
func WithInitRunnable(r PrefixRunnable) Option {
	return func(s *application) {
		if _, ok := s.initServices[r.GetPrefix()]; ok {
			log.Fatalf("prefix %s is duplicated", r.GetPrefix())
		}

		s.initServices[r.GetPrefix()] = r
	}
}

// Add sub service to SDK
func WithRunnable(r Runnable) Option {
	return func(s *application) { s.subServices = append(s.subServices, r) }
}
