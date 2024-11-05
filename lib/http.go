package lib

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type HttpHandler struct {
	Engine *fiber.App
}

func NewHttpHandler() *HttpHandler {
	engine := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	engine.Use(cors.New())
	engine.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	httpHandler := &HttpHandler{
		Engine: engine,
	}
	return httpHandler
}

func (fiberApp *HttpHandler) ShutDown() error {
	return fiberApp.Engine.Shutdown()
}
