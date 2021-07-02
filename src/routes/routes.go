package routes

import (
	"promoter/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	api.Post("/api/admin/register", controllers.Register)
}
