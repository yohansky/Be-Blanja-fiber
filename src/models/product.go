package models

import "gorm.io/gorm"

type Product struct {
	Id          uint `json:"id"`
	Name        string
	Rating      string
	Price       float64
	Color       string
	Size        string
	Amount      float64
	Condition   string
	Description string
	Image       string
	UserId      uint
	User        User `gorm:"foreignKey:UserId"`
	// CategoryId  uint
	// Category    Category `gorm:"foreignKey:CategoryId"`
}

func (product *Product) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Product{}).Count(&total)

	return total
}

func (product *Product) Take(db *gorm.DB, limit int, offset int) interface{} {
	var products []Product

	db.Offset(offset).Limit(limit).Find(&products)

	return products
}
