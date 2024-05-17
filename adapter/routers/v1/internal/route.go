package internal

import (
	"app/adapter/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, adapter *core.AdapterService) {
	/**
	 * @Summary Ping endpoint
	 * @Description ping pong
	 * @Tags ping
	 * @Accept json
	 * @Produce json
	 * @Success 200 {object} gin.H
	 * @Router /ping [get]
	 */
	r.GET("/ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"ping": "pong from v1/internal"}) })
}
