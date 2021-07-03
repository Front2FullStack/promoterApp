package routes

import (
	"promoter/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	api.Post("admin/register", controllers.Register)
	api.Post("admin/login", controllers.Login)
	api.Get("admin/user", controllers.User)
	api.Post("admin/logout", controllers.Logout)
}
