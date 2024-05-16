package models

import "gorm.io/gorm"

type Payment struct {
	Id           uint `json:"id"`
	UserId       uint
	User         User `gorm:"foreignKey:UserId"`
	BankId       uint
	Bank         Bank `gorm:"foreignKey:BankId"`
	TotalPayment string
	Status       string
}

func (payment *Payment) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Payment{}).Count(&total)

	return total
}

func (payment *Payment) Take(db *gorm.DB, limit int, offset int) interface{} {
	var payments []Payment

	db.Offset(offset).Limit(limit).Find(&payments)

	return payments
}
