package user_router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	user_structs "github.com/go-project/user/structs"
)

func SetRouter(app *gin.Engine) {
	app.POST("user", handleCreate)
	app.GET("user/:id", handleGet)
	app.GET("user", handleList)
	app.PATCH("user/:id", handlePatch)
	app.DELETE("user/:id", handleDelete)
}

func handleCreate(c *gin.Context) {
	payload := user_structs.CreateUser{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := create(payload)

	c.JSON(200, user)
}

func handleGet(c *gin.Context) {
	ID := c.Param("id")

	if IDint, err := strconv.Atoi(ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if IDint, err := strconv.Atoi(ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
	} else {
		user := patch(payload, IDint)

		c.JSON(200, user)
	}
}

func handleDelete(c *gin.Context) {
	ID := c.Param("id")

	if IDint, err := strconv.Atoi(ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
	} else {
		user := delete(IDint)

		c.JSON(200, user)
	}
}
