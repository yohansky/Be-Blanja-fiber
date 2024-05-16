package models

import "gorm.io/gorm"

type Bank struct {
	Id   uint `json:"id"`
	Name string
}

func (bank *Bank) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Bank{}).Count(&total)

	return total
}

func (bank *Bank) Take(db *gorm.DB, limit int, offset int) interface{} {
	var banks []Bank

	db.Offset(offset).Limit(limit).Find(&banks)

	return banks
}
