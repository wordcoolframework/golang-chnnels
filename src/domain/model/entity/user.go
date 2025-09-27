package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // customer, admin, seller
	Orders   []Order
	Products []Product `gorm:"foreignKey:UserID"`
}
