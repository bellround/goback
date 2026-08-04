package httptest

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

type SomeJSONInfo struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func jsontest(c fiber.Ctx) {
	data, err := json.Marshal(SomeJSONInfo{
		Message:   "Hello, World!",
		Timestamp: time.Now(),
	})

	if err != nil {
		log.Panic(err)
	}

	c.Send(data)
}
