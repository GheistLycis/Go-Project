package main

import (
	"github.com/gin-gonic/gin"
	user_router "github.com/go-project/user/router"
)

func main() {
	app := gin.Default()

	user_router.SetRouter(app)

	app.Run(":3000")
}
