package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hebobibun/go-ecommerce/handler"
)

var (
	user = handler.NewUser()
)

func UserRoutes(app *fiber.App) {
	app.Post("/user/register", user.RegisterUser)
	app.Post("/user/login", user.LoginUser)
}
