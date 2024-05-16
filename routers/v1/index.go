package router_v1

import (
	"app/routers/v1/external"
	"app/routers/v1/internal"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	internal.RegisterRoutes(router.Group("/internal"))
	external.RegisterRoutes(router.Group("/external"))
}
