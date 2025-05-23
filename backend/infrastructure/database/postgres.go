package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type postgres struct {
	client *pgx.Conn
}

func NewPostgresClient(username, password, url string, port int, dbName string, ssl bool) DatabaseService[*pgx.Conn] {
	sslMode := "disable"

	if ssl {
		sslMode = "enable"
	}

	db := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", username, password, url, port, dbName, sslMode)

	conn, err := pgx.Connect(context.Background(), db)
	if err != nil {
		panic(err)
	}

	return &postgres{conn}
}

func (db *postgres) GetInstance() *pgx.Conn {
	return db.client
}

func (db *postgres) Close() {
	db.client.Close(context.Background())
}
