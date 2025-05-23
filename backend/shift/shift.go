package shift

import (
	"payd/infrastructure/database"
	"payd/infrastructure/http"
	"payd/shift/handler"

	"github.com/jackc/pgx/v5"
)

func New(server http.HttpServer, db database.DatabaseService[*pgx.Conn]) {
	s := server.GetInstance()
	handler.ShiftRouter(s)
}
