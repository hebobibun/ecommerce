package interfaces

import "github.com/gofiber/fiber/v2"

type User interface {
	// User
	RegisterUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
}
