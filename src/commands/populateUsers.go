package main

import (
	"promoter/src/database"
	"promoter/src/models"

	"github.com/bxcodec/faker/v3"
)

func main() {
	database.Connect()
	for i := 0; i < 30; i++ {
		promoter := models.User{
			FirstName:  faker.FirstName(),
			LastName:   faker.LastName(),
			Email:      faker.Email(),
			IsPromoter: true,
		}

		promoter.SetPassword("1234")

		database.DB.Create(&promoter)
	}
}
