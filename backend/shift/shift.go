package shift

import (
	"payd/infrastructure/database"
	"payd/shift/handler"
	"payd/shift/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
)

func New(database database.DatabaseService) *fiber.App {
	r := repository.NewShiftsRepository(database)
	return handler.Routes(r)
}
