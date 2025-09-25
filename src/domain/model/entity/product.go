package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Stock       int
	UserID      int
	User        *User
	Categories  []Category `gorm:"many2many:product_categories"`
}

func (Product) TableName() string {
	return "products"
}
