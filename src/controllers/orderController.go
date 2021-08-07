package controllers

import (
	"fmt"
	"promoter/src/database"
	"promoter/src/models"

	"github.com/gofiber/fiber/v2"
)

func Orders(c *fiber.Ctx) error {

	var orders []models.Order

	database.DB.Preload("OrderItems").Find(&orders)

	for i, order := range orders {
		orders[i].Name = order.FullName()
		orders[i].Total = order.GetTotal()
	}
	return c.JSON(orders)
}

type CreateOrderRequest struct {
	Code      string
	FirstName string
	LastName  string
	Email     string
	Address   string
	Country   string
	City      string
	Zip       string
	Products  []map[string]int
}

func CreateOrder(c *fiber.Ctx) error {
	var request CreateOrderRequest

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	link := models.Link{
		Code: request.Code,
	}

	database.DB.Preload("User").First(&link)

	if link.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Link",
		})
	}

	order := models.Order{
		Code:          link.Code,
		UserId:        link.UserId,
		PromoterEmail: link.User.Email,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Email:         request.Email,
		Address:       request.Address,
		Country:       request.Country,
		City:          request.City,
		Zip:           request.Zip,
	}

	// Using transaction for creating order and order items
	tx := database.DB.Begin()
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fmt.Printf("Hello %v", request)

	for _, requestProduct := range request.Products {

		product := models.Product{}
		product.Id = uint(requestProduct["product_id"])
		database.DB.First(&product)

		total := product.Price * float64(requestProduct["quantity"])

		item := models.OrderItem{
			OrderId:         order.Id,
			ProductTitle:    product.Title,
			Price:           product.Price,
			Quantity:        uint(requestProduct["quantity"]),
			PromoterRevenue: 0.1 * total,
			AdminRevenue:    0.9 * total,
		}

		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

	}
	// to finish tx
	tx.Commit()
	return c.JSON(order)
}
