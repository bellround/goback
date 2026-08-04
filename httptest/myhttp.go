package httptest

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	app *fiber.App
}

func greet(c fiber.Ctx) {
	c.SendString(fmt.Sprintf("Hello World! %s", time.Now()))
}

func New() *Server {
	app := fiber.New()

	app.Get("/", greet)
	app.Get("/json", jsontest)

	return &Server{app: app}
}

func (s *Server) Run() error {
	return s.app.Listen(":8080")
}

func (s *Server) Stop(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}
