package models

import "gorm.io/gorm"

type Category struct {
	Id    uint `json:"id"`
	Name  string
	Image string
}

func (category *Category) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Category{}).Count(&total)

	return total
}

func (category *Category) Take(db *gorm.DB, limit int, offset int) interface{} {
	var addresses []Category

	db.Offset(offset).Limit(limit).Find(&addresses)

	return addresses
}
