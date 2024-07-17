package interfaces

import "github.com/gofiber/fiber/v2"

type Admin interface {
	Create(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}
