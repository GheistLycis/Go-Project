package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	service "github.com/go-project/app/user/service"
)

func delete(c *gin.Context) {
	ID := c.Param("id")

	if IDint, err := strconv.Atoi(ID); err != nil {
		c.JSON(http.StatusBadRequest, "ID inválido")
	} else {
		user := service.Delete(IDint)

		c.JSON(http.StatusOK, user)
	}
}
