package main

import (
	"payd/application"
	"payd/config"
)

func main() {
	c := config.Get()

	app := application.New(c)
	app.Run()
}
