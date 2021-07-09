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
