package server

import (
	"app/adapter/routers/middleware"
	"app/addons"
	"app/addons/logger"
	"app/gosdk"
	"app/utils"
	"context"
	"flag"
	"net/http"
	"time"

	_ "app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerGinOpt struct {
	Prefix addons.AddOnPrefix
	Port   int
	Mode   string
}

type GinServer struct {
	engine     *gin.Engine
	httpServer *http.Server
}

type server struct {
	*ServerGinOpt
	serverCtx      *GinServer
	applicationCtx gosdk.ApplicationContext
	logger         logger.Logger
	name           addons.AddOnName
}

func NewGinServer(name addons.AddOnName, prefix addons.AddOnPrefix) *server {
	return &server{
		name: name,
		ServerGinOpt: &ServerGinOpt{
			Prefix: prefix,
		},
	}
}

// Implement gosdk.Application interface
func (s *server) Name() string { return string(s.name) }

func (s *server) GetPrefix() string { return string(s.Prefix) }

func (s *server) Get() interface{} { return s.serverCtx }

// Implement gosdk.PrefixRunnable
func (s *server) InitFlags() {
	prefix := string(s.GetPrefix())
	if prefix != "" {
		prefix += "-"
	}

	flag.IntVar(&s.Port, prefix+"port", 8402, "Port to listen on")
	flag.StringVar(&s.Mode, "gin-mode", gin.ReleaseMode, "Gin mode")
}

func (s *server) Configure() error {
	s.logger = logger.GetCurrent().GetLogger(string(s.name))
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middleware.WithCors())

	// s.serverCtx.engine = engine
	// s.serverCtx.httpServer = &http.Server{
	// 	Addr:    ":" + utils.Int2String(s.Port),
	// 	Handler: s.serverCtx.engine,
	// }

	s.serverCtx = &GinServer{
		engine: engine,
		httpServer: &http.Server{
			Addr: ":" + utils.Int2String(s.Port),
		},
	}

	return nil
}

func (s *server) Run() error {
	if err := s.Configure(); err != nil {
		return err
	}
	s.logger.Info("Starting server")
	return nil
}

func (s *server) Stop() <-chan bool {
	s.logger.Infof("Stopping %s", s.Name())
	c := make(chan bool)

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Attempt to gracefully shutdown the server
		if err := s.serverCtx.httpServer.Shutdown(ctx); err != nil {
			s.logger.Errorf("Server forced to shutdown: %v", err)
		} else {
			s.logger.Infof("Server %s stopped gracefully", s.Name())
		}

		c <- true
	}()
	return c
}

type RouteOption func(*server)
type RouterHandler func(*gin.RouterGroup)

func (s *GinServer) Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return s.engine.Group(relativePath, handlers...)
}

func (s *GinServer) ListenAndServe() error {
	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.httpServer.Handler = s.engine
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
