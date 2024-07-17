package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hebobibun/go-ecommerce/handler"
)

var (
	admin = handler.NewAdmin()
)

func AdminRoutes(app *fiber.App) {
	app.Post("/admin/create", admin.Create)
	app.Post("/admin/login", admin.Login)
}
