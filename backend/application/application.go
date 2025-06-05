package application

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"payd/infrastructure/database"
	"payd/infrastructure/http/server"
	"payd/shift"
)

type Application struct {
	Database database.DatabaseService
	Server   http.HttpServer
}

func New() *Application {
	db := database.NewPostgres("admin", "admin", "localhost", 5442, "payd", false)
	return &Application{
		Database: db,
		Server:   http.NewServer(9095),
	}
}

func (a *Application) Run() {
	a.start()

	<-a.wait()

	stopCtx, stopCancellation := context.WithTimeout(context.Background(), time.Second*20)
	defer stopCancellation()
	a.stop(stopCtx)
}

func (a *Application) start() {
	s := shift.New(a.Database)
	a.Server.AddRoute("/shifts", s)
	go a.Server.Start()
}

func (a *Application) wait() <-chan os.Signal {
	// Create channel to signify a signal being sent
	c := make(chan os.Signal, 1)

	// When an interrupt or termination signal is sent, notify the channel
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	return c
}

func (a *Application) stop(ctx context.Context) {
	a.Server.Shutdown(ctx)
}
