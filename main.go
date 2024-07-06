package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file for local development
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		name := os.Getenv("NAME")
		if name == "" {
			name = "World"
		}
		return c.SendString("Hello " + name)
	})

	log.Fatal(app.Listen(":3000"))
}
