package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	service "github.com/go-project/app/user/service"
	structs "github.com/go-project/app/user/structs"
)

func post(c *gin.Context) {
	payload := structs.CreateUser{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			fieldErrors := make(map[string]string)

			for _, e := range errors {
				fieldErrors[e.Field()] = e.Tag()
			}

			c.JSON(http.StatusBadRequest, fieldErrors)
			return
		}

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user := service.Create(payload)

	c.JSON(200, user)
}
