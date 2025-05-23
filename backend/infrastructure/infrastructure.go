package infrastructure

import (
	"payd/infrastructure/database"
	"payd/infrastructure/http"

	"github.com/jackc/pgx/v5"
)

type Infrastructure struct {
	Database   database.DatabaseService[*pgx.Conn]
	HttpServer http.HttpServer
}

func New() *Infrastructure {
	db := database.NewPostgresClient("admin", "admin", "localhost", 5442, "payd", false)
	server := http.NewServer(9095)
	return &Infrastructure{
		Database:   db,
		HttpServer: server,
	}
}

func (i *Infrastructure) RunWebServer() {
	i.HttpServer.Run()
}

func (i *Infrastructure) CloseDatabase() {
	i.Database.Close()
}
