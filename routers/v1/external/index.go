package external

import (
	controller "gin-seed/adapter/controller/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"ping": "pong from v1/external cc"}) })
  r.GET("/users", controller.ListUsers)
}
