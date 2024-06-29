package user_controller

import (
	"github.com/gin-gonic/gin"
)

func SetRouter(server *gin.Engine) {
	server.POST("user", handleCreate)
	server.GET("user/:id", handleGet)
	server.GET("user", handleList)
	server.PATCH("user/:id", handlePatch)
	server.DELETE("user/:id", handleDelete)
}
