package user_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	user_service "github.com/go-project/app/user/service"
	user_structs "github.com/go-project/app/user/structs"
)

func patch(c *gin.Context) {
	payload := user_structs.PatchUser{}
	ID := c.Param("id")

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

	if IDint, err := strconv.Atoi(ID); err != nil {
		c.JSON(http.StatusBadRequest, "ID inválido")
	} else {
		user := user_service.Patch(payload, IDint)

		c.JSON(200, user)
	}
}
