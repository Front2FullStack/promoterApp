package routes

import (
	"promoter/src/controllers"
	"promoter/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")
	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)

	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Get("user", controllers.User)
	adminAuthenticated.Post("logout", controllers.Logout)
	adminAuthenticated.Put("users/info", controllers.UpdateInfo)
	adminAuthenticated.Put("users/password", controllers.UpdatePassword)
	adminAuthenticated.Get("promoters", controllers.Promoters)
	adminAuthenticated.Get("products", controllers.Products)
	adminAuthenticated.Get("products/:id", controllers.GetProduct)
	adminAuthenticated.Put("products/:id", controllers.UpdateProduct)
	adminAuthenticated.Post("products", controllers.CreateProducts)
	adminAuthenticated.Delete("products/:id", controllers.DeleteProduct)
	adminAuthenticated.Get("users/:id/links", controllers.Link)
	adminAuthenticated.Get("orders", controllers.Orders)

	promoter := api.Group("promoter")
	promoter.Post("register", controllers.Register)
	promoter.Post("login", controllers.Login)
	promoter.Get("products/frontend", controllers.ProductsFrontend)

	promoterAuthenticated := promoter.Use(middlewares.IsAuthenticated)
	promoterAuthenticated.Get("user", controllers.User)
	promoterAuthenticated.Post("logout", controllers.Logout)
	promoterAuthenticated.Put("users/info", controllers.UpdateInfo)
	promoterAuthenticated.Put("users/password", controllers.UpdatePassword)
}
