package shift

import (
	"payd/infrastructure"
	"payd/shift/handler"
	"payd/shift/infrastructure/repository"
)

func New(infra *infrastructure.Infrastructure) {
	server := infra.HttpServer.GetInstance()
	database := infra.Database
	r := repository.NewShiftsRepository(database)
	handler.ShiftRouter(server, r)
}
