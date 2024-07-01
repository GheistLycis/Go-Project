package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-project/app/shared"
	"github.com/go-project/app/user/service"
	"github.com/go-project/app/user/structs"
)

func patch(c *gin.Context) {
	ID := c.Param("id")
	payload := structs.PatchUser{}

	IDint, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID inválido")
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			shared.HandleValidationError(errors, c)
			return
		}

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := service.Patch(payload, IDint)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID inválido")
		return
	}

	c.JSON(http.StatusOK, user)
}
