package controller

import (
	"github.com/gin-gonic/gin"
)

func SetRouter(server *gin.Engine) {
	server.POST("user", post)
	server.GET("user/:id", get)
	server.GET("user", list)
	server.PATCH("user/:id", patch)
	server.DELETE("user/:id", delete)
}
