package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hebobibun/go-ecommerce/handler"
	"github.com/hebobibun/go-ecommerce/middleware"
)

var (
	admin   = handler.NewAdmin()
	product = handler.NewProduct()
)

func AdminRoutes(app *fiber.App) {
	app.Post("/admin/create", admin.Create)
	app.Post("/admin/login", admin.Login)

	admin := app.Group("/admin", middleware.JWT)

	admin.Post("/categories", product.CreateCategory)
	admin.Post("/products", product.CreateProduct)
}
