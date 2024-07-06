package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/faridlan/vercel-go/service"
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

	app.Get("/profile", func(c *fiber.Ctx) error {

		profile, err := service.ConvertJson()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error reading profile data")
		}
		return c.SendString(fmt.Sprintf("Username: %s\nEmail: %s", profile.Username, profile.Email))
	})

	adaptor.FiberApp(app)(w, r)
}
