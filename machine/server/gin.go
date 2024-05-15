package server

import (
	"flag"
	"gin-seed/gosdk"
	"gin-seed/machine/logger"
	"gin-seed/routers/middleware"
	router_v1 "gin-seed/routers/v1"
	"gin-seed/utils"

	"github.com/gin-gonic/gin"
)

type ServerGinOpt struct {
	Prefix string
	Port   int
}

type server struct {
	*ServerGinOpt
	applicationCtx gosdk.ApplicationContext
	logger         logger.Logger
	name           string
	router         *gin.Engine
}

func NewGinServer(name string, prefix string) *server {
	return &server{
		name: name,
		ServerGinOpt: &ServerGinOpt{
			Prefix: prefix,
		},
	}
}

// Implement gosdk.Application interface
func (s *server) Name() string { return s.name }

func (s *server) GetPrefix() string { return s.Prefix }

func (s *server) Get() interface{} { return s }

// Implement gosdk.PrefixRunnable
func (s *server) InitFlags() {
	flag.IntVar(&s.Port, "PORT", 8402, "Port to listen on")
}

func (s *server) Configure() error {
	s.logger = logger.GetCurrent().GetLogger(s.name)
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.WithCors())

	// Register all routes
	router_v1.RegisterRoutes(router.Group("/v1"))

	s.router = router

	return nil
}

func (s *server) Run() error {
	s.Configure()
	s.logger.Info("Starting server")
	s.router.Run(":" + utils.Int2String(s.Port))
	return nil
}

func (s *server) Stop() <-chan bool {
	c := make(chan bool)
	go func() {
		c <- true
		s.logger.Info("Server stopped")
	}()
	return c
}
