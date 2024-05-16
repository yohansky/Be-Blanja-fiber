package helper

import (
	"backend-gin/src/config"
	"backend-gin/src/models"
)

func Migrate() {
	config.DB.AutoMigrate(&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Product{},
		&models.Custommer{},
		&models.Seller{},
		&models.Bank{},
		&models.Order{},
		&models.Payment{},
		&models.Address{},
		&models.Delivery{})
}
