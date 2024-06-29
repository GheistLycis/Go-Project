package user_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	user_service "github.com/go-project/app/user/service"
)

func get(c *gin.Context) {
	ID := c.Param("id")

	if IDint, err := strconv.Atoi(ID); err != nil {
		c.JSON(http.StatusBadRequest, "ID inválido")
	} else {
		user := user_service.Get(IDint)

		c.JSON(200, user)
	}
}

func list(c *gin.Context) {
	users := user_service.List()

	c.JSON(200, users)
}
