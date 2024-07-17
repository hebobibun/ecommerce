package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	database "github.com/hebobibun/go-ecommerce/db"
	"github.com/hebobibun/go-ecommerce/routes"
)

func init() {
	database.Migrate()
	database.InitRedis()
}

func main() {

	app := fiber.New()
	routes.AdminRoutes(app)

	log.Println("Server running on port " + os.Getenv("PORT"))

	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
