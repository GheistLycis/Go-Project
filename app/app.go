package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_controller "github.com/go-project/app/user/controller"
)

var server *gin.Engine

func initRouters() {
	user_controller.SetRouter(server)
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
