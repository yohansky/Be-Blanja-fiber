package models

import "gorm.io/gorm"

type Order struct {
	Id          uint `json:"id"`
	Order_Size  string
	Order_Color string
	Quantity    uint
	Total_Price uint
	SellerId    uint
	UserId      uint
	ProductId   uint
	Product     Product `gorm:"foreignKey:ProductId"`
	Status      string
}

func (order *Order) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Order{}).Count(&total)

	return total
}

func (order *Order) Take(db *gorm.DB, limit int, offset int) interface{} {
	var orders []Order

	db.Offset(offset).Limit(limit).Find(&orders)

	return orders
}

func (o *Order) BeforeSave(tx *gorm.DB) (err error) {
	// Set default value only if the Status field is empty
	if o.Status == "" {
		o.Status = "pending"
	}
	return
}
