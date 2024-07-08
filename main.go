package main

import (
	"fmt"
	"log"
	"os"

	"github.com/faridlan/vercel-go/service"
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

	app.Get("/profile", func(c *fiber.Ctx) error {

		profile, err := service.ConvertJson()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error reading profile data: %v", err))
		}
		return c.SendString(fmt.Sprintf("Username: %s\nEmail: %s", profile.Username, profile.Email))

	})

	app.Get("/wd", func(c *fiber.Ctx) error {

		path, _ := os.Getwd()
		return c.SendString(path)

	})

	app.Get("/dir", func(c *fiber.Ctx) error {

		directory := []string{}
		path, _ := os.Getwd()
		dirs, _ := os.ReadDir(path)

		for _, dir := range dirs {
			directory = append(directory, dir.Name())
		}

		return c.JSON(directory)

	})

	log.Fatal(app.Listen(":3000"))
}
