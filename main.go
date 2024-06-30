package main

import (
	"github.com/go-project/app"
	"github.com/go-project/database"
)

func init() {
	database.Init(true)
}

func main() {
	app.Init(3000)
}
