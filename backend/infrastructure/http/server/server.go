package http

import (
	"context"
	"errors"
	"fmt"
	http1 "net/http"
	"time"

	"payd/infrastructure/http/server/middleware"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type routeFn func() (string, *fiber.App)

type HttpServer interface {
	AddRoutes(...routeFn)
	AddMiddleware(...fiber.Handler)
	Start() error
	Shutdown(context.Context) error
}

type server struct {
	instance *fiber.App
	port     int
}

func NewServer(port int) HttpServer {
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	s := &server{
		instance: app,
		port:     port,
	}

	s.AddMiddleware(
		cors.New(),
		logger.New(),
		middleware.ErrorHandler,
		otelfiber.Middleware(),
	)

	return s
}

func (s *server) AddMiddleware(fn ...fiber.Handler) {
	for _, f := range fn {
		s.instance.Use(f)
	}
}

func (s *server) AddRoutes(fn ...routeFn) {
	for _, f := range fn {
		name, router := f()
		s.instance.Mount(name, router)
	}
}

func (s *server) Start() error {
	if err := s.instance.Listen(fmt.Sprintf(":%d", s.port)); err != nil && !errors.Is(err, http1.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *server) Shutdown(ctx context.Context) error {
	if err := s.instance.ShutdownWithContext(ctx); err != nil {
		return err
	}

	return nil
}
