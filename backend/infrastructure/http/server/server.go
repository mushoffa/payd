package http

import (
	"context"
	"errors"
	"fmt"
	http1 "net/http"
	"time"

	"payd/infrastructure/http/server/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type HttpServer interface {
	AddRoute(string, *fiber.App)
	Start() error
	Shutdown(context.Context) error
}

type server struct {
	instance *fiber.App
	port     int
	routes   router
}

type router struct {
	v1 fiber.Router
	// Extend the version api router here:
	//   v2 fiber.Router
	// Add new router(s) accordingly at instantiation function
}

func NewServer(port int) HttpServer {
	app := fiber.New(fiber.Config{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	// Middleware
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(middleware.ErrorHandler)

	// Router
	api := app.Group("/api")
	v1 := api.Group("/v1")

	r := router{
		v1: v1,
	}

	return &server{
		instance: app,
		port:     port,
		routes:   r,
	}
}

func (s *server) AddRoute(name string, router *fiber.App) {
	s.routes.v1.Mount(name, router)
}

func (s *server) Start() error {
	if err := s.instance.Listen(fmt.Sprintf(":%d", s.port)); err != nil && !errors.Is(err, http1.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *server) Shutdown(ctx context.Context) error {
	if err := s.instance.ShutdownWithContext(ctx); err != nil {
		fmt.Println("Err: ", err)
		return err
	}

	return nil
}
