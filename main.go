package main

import (
	"promoter/src/database"
	"promoter/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
