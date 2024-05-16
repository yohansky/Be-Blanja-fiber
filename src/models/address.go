package models

import "gorm.io/gorm"

type Address struct {
	Id             uint `json:"id"`
	Recipient_Name string
	Address_as     string
	Address        string
	Phone          string
	Postal_Code    string
	City           string
	UserId         uint
	User           User `gorm:"foreignKey:UserId"`
}

func (address *Address) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Address{}).Count(&total)

	return total
}

func (address *Address) Take(db *gorm.DB, limit int, offset int) interface{} {
	var addresses []Address

	db.Offset(offset).Limit(limit).Find(&addresses)

	return addresses
}
