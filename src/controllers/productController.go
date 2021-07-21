package controllers

import (
	"context"
	"encoding/json"
	"promoter/src/database"
	"promoter/src/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Products(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)
	return c.JSON(products)
}

func CreateProducts(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	var product models.Product

	id, _ := strconv.Atoi(c.Params("id"))

	product.Id = uint(id)

	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(&product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&product)

	return nil
}

func ProductsFrontend(c *fiber.Ctx) error {
	var products []models.Product
	var ctx = context.Background()
	// fetching result from redis
	result, err := database.Cache.Get(ctx, "products_frontend").Result()
	if err != nil {
		database.DB.Find(&products)
		// converting db results into bytes
		bytes, err := json.Marshal(products)
		if err != nil {
			panic(err)
		}
		// validating and checking for err while setting up cache
		if errKey := database.Cache.Set(ctx, "products_frontend", bytes, 30*time.Minute).Err(); errKey != nil {
			panic(errKey)
		}
	} else {
		// converting redis result into json and casting
		json.Unmarshal([]byte(result), &products)
	}

	return c.JSON(products)
}
