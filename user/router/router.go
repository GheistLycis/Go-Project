package user_router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_structs "github.com/go-project/user/structs"
)

func SetRouter(app *gin.Engine) {
	app.GET("user/:id", func(c *gin.Context) {
		ID := c.Param("id")

		user := Get(ID)

		c.JSON(200, gin.H{"user": user})
	})

	app.GET("user", func(c *gin.Context) {
		users := List()

		c.JSON(200, gin.H{"users": users})
	})

	app.POST("user", func(c *gin.Context) {
		var payload user_structs.CreateUser

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		user := Create(payload)

		c.JSON(200, gin.H{"user": user})
	})
}
