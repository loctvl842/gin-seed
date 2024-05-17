package external

import (
	controller "app/adapter/controller/v1"
	"app/adapter/core"
	router "app/adapter/routers"
	"app/adapter/service"
	"app/addons"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, as *core.AdapterService) {
	r.GET("/ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"ping": "pong from v1/external"}) })

	db := as.AppCtx.MustGet(string(addons.PgDatabasePrefix)).(*gorm.DB)
	userSv := service.NewUserService(db)
	userController := controller.NewUserController(userSv)
	r.POST("/users", router.NewHandler(as, userController.CreateUser))
}
