package controllers

import (
	"promoter/src/database"
	"promoter/src/models"

	"github.com/gofiber/fiber/v2"
)

func Promoters(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Where("is_promoter = true").Find(&users)
	return c.JSON(users)
}

func Rankings(c *fiber.Ctx) error {

	var users []models.User

	database.DB.Find(&users, models.User{
		IsPromoter: true,
	})

	var result []interface{}

	for _, user := range users {
		promoter := models.Promoter(user)
		promoter.CalculateRevenue(database.DB)

		result = append(result, fiber.Map{
			user.Name(): promoter.Revenue,
		})
	}

	return c.JSON(result)

}
