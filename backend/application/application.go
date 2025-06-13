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
	shutdown, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP)
	defer stop()

	a.start()

	<-shutdown.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	a.stop(ctx)
	os.Exit(0)
}

func (a *Application) start() {
	shift_feat := shift.New(a.Database)
	a.Server.AddRoutes(
		shift_feat.Handler().RegisterV1,
		shift_feat.Handler().RegisterV2,
	)

	go a.Server.Start()
}

func (a *Application) stop(ctx context.Context) {
	a.Server.Shutdown(ctx)
	a.Database.Close()
	a.Tracer.Shutdown(ctx)
}
