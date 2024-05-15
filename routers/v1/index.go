package router_v1

import (
	"gin-seed/routers/v1/external"
	"gin-seed/routers/v1/internal"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	internal.RegisterRoutes(router.Group("/internal"))
	external.RegisterRoutes(router.Group("/external"))
}
