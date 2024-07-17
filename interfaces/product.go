package interfaces

import "github.com/gofiber/fiber/v2"

type Product interface {
	// Category
	CreateCategory(c *fiber.Ctx) error

	// Product
	CreateProduct(c *fiber.Ctx) error
}
