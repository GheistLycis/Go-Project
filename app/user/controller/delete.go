package user_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	user_service "github.com/go-project/app/user/service"
)

func handleDelete(c *gin.Context) {
	ID := c.Param("id")

	if IDint, err := strconv.Atoi(ID); err != nil {
		c.JSON(http.StatusBadRequest, "ID inválido")
	} else {
		user := user_service.Delete(IDint)

		c.JSON(200, user)
	}
}
