package http

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type HttpServer interface {
	GetInstance() *fiber.App
	Run()
}

type server struct {
	instance *fiber.App
	port     int
}

func NewServer(port int) HttpServer {
	app := fiber.New()
	app.Use(logger.New())

	s := &server{app, port}
	// s.handler(u)

	return s
}

func (s *server) GetInstance() *fiber.App {
	return s.instance
}

func (s *server) Run() {
	c := make(chan os.Signal, 1) // Create channel to signify a signal being sent

	go func() {
		if err := s.instance.Listen(fmt.Sprintf(":%d", s.port)); err != nil {
			log.Panic(err)
		}
	}()

	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	_ = <-c                                         // This blocks the main thread until an interrupt is received
	_ = s.instance.Shutdown()

	fmt.Println("HTTP Server was successfully shutdown.")
}
