package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id         uint     `json:"id"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Email      string   `json:"email" gorm:"unique"`
	Password   []byte   `json:"-"`
	IsPromoter bool     `json:"-"`
	Revenue    *float64 `json:"revenue,omitempty" gorm:"-"`
}

//Aliases
type Admin User
type Promoter User

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

func (user *User) Name() string {
	return user.FirstName + "" + user.LastName
}

func (admin *Admin) CalculateRevenue(db *gorm.DB) {
	var orders []Order

	db.Preload("OrderItems").Find(&orders, &Order{
		UserId:   admin.Id,
		Complete: true,
	})

	var revenue float64 = 0

	for _, order := range orders {
		for _, orderItem := range order.OrderItems {
			revenue += orderItem.AdminRevenue
		}
	}
	admin.Revenue = &revenue

}

func (promoter *Promoter) CalculateRevenue(db *gorm.DB) {
	var orders []Order

	db.Preload("OrderItems").Find(&orders, &Order{
		UserId:   promoter.Id,
		Complete: true,
	})

	var revenue float64 = 0

	for _, order := range orders {
		for _, orderItem := range order.OrderItems {
			revenue += orderItem.PromoterRevenue
		}
	}
	promoter.Revenue = &revenue

}
