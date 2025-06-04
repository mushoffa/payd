package shift

import (
	"payd/infrastructure/database"
	"payd/shift/handler"
	"payd/shift/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func New(database database.DatabaseService[*pgx.Conn]) *fiber.App {
	r := repository.NewShiftsRepository(database)
	return handler.Routes(r)
}
