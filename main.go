package main

import (
	"github.com/gin-gonic/gin"
	user_router "github.com/go-project/user/router"
)

var app = gin.Default()

func init() {
	user_router.SetRouter(app)
}

func main() {
	app.Run(":3000")
}
