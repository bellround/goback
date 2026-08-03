package httptest

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
)

func greet(c fiber.Ctx) {
	c.SendString(fmt.Sprintf("Hello World! %s", time.Now()))
}

func RunHTTP() error {
	app := fiber.New()

	app.Get("/", greet)

	return app.Listen(":8080")
}
