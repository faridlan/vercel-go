package api

import (
	"net/http"
	"os"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		name := os.Getenv("NAME")
		if name == "" {
			name = "World"
		}
		return c.SendString("Hello " + name)
	})

	adaptor.FiberApp(app)(w, r)
}
