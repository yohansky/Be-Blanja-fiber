package models

import "gorm.io/gorm"

type Delivery struct {
	Id    uint `json:"id"`
	Price uint
}

func (delivery *Delivery) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Delivery{}).Count(&total)

	return total
}

func (delivery *Delivery) Take(db *gorm.DB, limit int, offset int) interface{} {
	var deliverys []Delivery

	db.Offset(offset).Limit(limit).Find(&deliverys)

	return deliverys
}

// func (D *Delivery) BeforeSave(tx *gorm.DB) (err error) {
// 	if D.Price == 0 {
// 		D.Price = 10000
// 	}
// 	return
// }
