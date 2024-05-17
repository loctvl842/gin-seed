package v1

import (
	"app/adapter/core"
	"app/adapter/model"
	"app/adapter/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *service.UserService
}

func NewUserController(s *service.UserService) *UserController {
	return &UserController{Service: s}
}

// ListUsers is a sample handler function.
// @Summary List users
// @Description Create Users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} string
// @Router /v1/external/users [post]
func (uc *UserController) CreateUser(s *core.AdapterService, c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.Service.UserRepository.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
