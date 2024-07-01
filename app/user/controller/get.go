package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	service "github.com/go-project/app/user/service"
)

func get(c *gin.Context) {
	ID := c.Param("id")
	IDint, err := strconv.Atoi(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, "ID inválido")
		return
	}

	user, err := service.Get(IDint)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func list(c *gin.Context) {
	users, err := service.List()

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, users)
}
