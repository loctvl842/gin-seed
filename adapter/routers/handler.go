package router

import (
	"app/adapter/core"

	"github.com/gin-gonic/gin"
)

/**
 * @Summary Custom HTTP handler function
 * @Description Execute the handler function with the given service and context
 * @Accept json
 * @Produce json
 * @Success 200 {object} gin.H
 */
type Handler func(service *core.AdapterService, ctx *gin.Context)

func NewHandler(service *core.AdapterService, handler Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler(service, ctx)
	}
}
