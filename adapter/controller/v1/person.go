package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, []string{"user1", "user2"})
}
