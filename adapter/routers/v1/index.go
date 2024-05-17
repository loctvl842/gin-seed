package router_v1

import (
	"app/adapter/routers/v1/external"
	"app/adapter/routers/v1/internal"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	internal.RegisterRoutes(router.Group("/internal"))
	external.RegisterRoutes(router.Group("/external"))
}
