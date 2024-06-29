package main

import (
	"github.com/go-project/app"
	"github.com/go-project/database"
)

func init() {
	database.Init()
	app.Setup()
}

func main() {
	app.Init()
}
