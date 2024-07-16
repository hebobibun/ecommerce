package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!\n")
	})

	log.Println("Server running on port " + os.Getenv("PORT"))

	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
