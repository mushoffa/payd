package shift

import (
	"payd/infrastructure/http"
	"payd/shift/handler"
)

func New(server http.HttpServer) {
	s := server.GetInstance()
	handler.ShiftRouter(s)
}
