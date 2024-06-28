package user_router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	user_structs "github.com/go-project/user/structs"
)

func SetRouter(server *gin.Engine) {
	server.POST("user", handleCreate)
	server.GET("user/:id", handleGet)
	server.GET("user", handleList)
	server.PATCH("user/:id", handlePatch)
	server.DELETE("user/:id", handleDelete)
}

func handleCreate(c *gin.Context) {
	payload := user_structs.CreateUser{}

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

	user := create(payload)

	c.JSON(200, user)
}

func handleGet(c *gin.Context) {
	ID := c.Param("id")

	if IDint, err := strconv.Atoi(ID); err != nil {
		c.JSON(http.StatusBadRequest, "ID inválido")
	} else {
		user := get(IDint)

		c.JSON(200, user)
	}
}

func handleList(c *gin.Context) {
	users := list()

	c.JSON(200, users)
}

func handlePatch(c *gin.Context) {
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
		user := patch(payload, IDint)

		c.JSON(200, user)
	}
}

func handleDelete(c *gin.Context) {
	ID := c.Param("id")

	if IDint, err := strconv.Atoi(ID); err != nil {
		c.JSON(http.StatusBadRequest, "ID inválido")
	} else {
		user := delete(IDint)

		c.JSON(200, user)
	}
}
