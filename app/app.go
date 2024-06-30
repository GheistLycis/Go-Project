package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	usercontroller "github.com/go-project/app/user/controller"
)

/*
Init creates the server, set its routers and handlers and then runs it.

-port: the server port.
*/
func Init(port int) {
	server := gin.Default()

	initRouters(server)
	initHandlers(server)
	server.Run(fmt.Sprintf(":%d", port))
}

func initRouters(server *gin.Engine) {
	usercontroller.SetRouter(server)
}

func initHandlers(server *gin.Engine) {
	server.Use(handleErrors)
}

func handleErrors(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors[0].Err

		c.JSON(http.StatusInternalServerError, err.Error())
	}
}
