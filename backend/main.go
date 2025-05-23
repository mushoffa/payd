package main

import (
	"fmt"

	// "payd/application"
	"payd/infrastructure/database"
	"payd/infrastructure/http"
	// "payd/infrastructure/repository"
	"payd/shift"
)

func main() {
	database.NewPostgresClient("admin", "admin", "localhost", 5442, "payd", false)
	// shift_r := repository.NewShiftRepository(db)

	// usecase := application.NewUsecase(shift_r)
	server := http.NewServer(9095)

	shift.New(server)
	server.Run()

	fmt.Println("Close database here")
	// db.Close()
}
