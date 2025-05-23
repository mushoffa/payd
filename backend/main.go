package main

import (
	"payd/infrastructure/database"
	"payd/infrastructure/http"
	"payd/shift"
)

func main() {
	postgres := database.NewPostgresClient("admin", "admin", "localhost", 5442, "payd", false)
	server := http.NewServer(9095)

	shift.New(server, postgres)
	server.Run()

	postgres.Close()
}
