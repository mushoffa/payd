package main

import (
	"payd/infrastructure"
	"payd/shift"
)

func main() {
	i := infrastructure.New()

	shift.New(i)
	i.RunWebServer()
	i.CloseDatabase()
}
