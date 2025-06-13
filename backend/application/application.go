package application

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"payd/config"
	"payd/infrastructure/database"
	"payd/infrastructure/http/server"
	"payd/infrastructure/trace"
	"payd/shift"
)

type Application struct {
	Database database.DatabaseService
	Server   http.HttpServer
	Tracer   trace.TracerService
}

func New(config *config.Config) *Application {
	db := database.NewPostgres(config)
	tracer := trace.NewOpenTelemetry(config)

	return &Application{
		Database: db,
		Server:   http.NewServer(9095),
		Tracer:   tracer,
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
	shift_feat := shift.New(a.Database)
	a.Server.AddRoutes(
		shift_feat.Handler().RegisterV1,
		shift_feat.Handler().RegisterV2,
	)

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
