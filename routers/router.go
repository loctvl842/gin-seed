package routers

import (
	"gin-seed/routers/middleware"
	router_v1 "gin-seed/routers/v1"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.WithCors())

  // Register all routes
	router_v1.RegisterRoutes(router.Group("/v1"))

	return router
}
