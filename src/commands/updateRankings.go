package main

import (
	"context"
	"promoter/src/database"
	"promoter/src/models"

	"github.com/go-redis/redis/v8"
)

func main() {
	database.Connect()
	database.SetupRedis()
	ctx := context.Background()
	var users []models.User

	database.DB.Find(&users, models.User{
		IsPromoter: true,
	})

	for _, user := range users {
		promoter := models.Promoter(user)
		promoter.CalculateRevenue(database.DB)

		database.Cache.ZAdd(ctx, "rankings", &redis.Z{
			Member: user.Name(),
			Score:  *promoter.Revenue,
		})
	}
}
