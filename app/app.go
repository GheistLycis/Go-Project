package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_router "github.com/go-project/user/router"
)

var server *gin.Engine

func initRouters() {
	user_router.SetRouter(server)
}

func handleErrors(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors[0].Err

		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func Setup() {
	server = gin.Default()

	initRouters()

	server.Use(handleErrors)
}

func Init() {

	server.Run(":3000")
}
