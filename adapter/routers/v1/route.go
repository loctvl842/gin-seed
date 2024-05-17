package router_v1

import (
	"app/adapter/core"
	"app/adapter/routers/v1/external"
	"app/adapter/routers/v1/internal"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, service *core.AdapterService) {
	internal.RegisterRoutes(router.Group("/internal"), service)
	external.RegisterRoutes(router.Group("/external"), service)
}
