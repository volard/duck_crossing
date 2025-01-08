package main

import (
	"github.com/gofiber/fiber/v2"
)

// HelloHandler handles the hello world endpoint
type HelloHandler struct{}

// NewHelloHandler creates a new HelloHandler
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

// RegisterRoutes registers the hello world route
func (h *HelloHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/hello", h.HelloWorld)
}

// HelloWorld responds with "Hello, World!"
func (h *HelloHandler) HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
