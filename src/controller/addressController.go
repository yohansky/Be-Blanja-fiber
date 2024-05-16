package controller

import (
	"backend-gin/src/config"
	"backend-gin/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllAddress(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Address{}, page))
}

func GetAddressByUserID(c *fiber.Ctx) error {
	id := c.Params("id")

	var address []models.Address
	if err := config.DB.Where("user_id = ?", id).Find(&address).Error; err != nil {
		return c.JSON(fiber.Map{"Error": "User not found"})
	}

	return c.JSON(address)
}

func CreateAddress(c *fiber.Ctx) error {
	var address models.Address

	if err := c.BodyParser(&address); err != nil {
		return err
	}

	config.DB.Create(&address)

	return c.JSON(address)
}

func GetAddress(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var address models.Address

	address.Id = uint(id)

	config.DB.Preload("User").Find(&address)

	return c.JSON(address)
}

func UpdateAddress(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var address models.Address

	address.Id = uint(id)

	if err := c.BodyParser(&address); err != nil {
		return err
	}

	config.DB.Model(&address).Updates(address)

	return c.JSON(address)
}

func DeleteAddress(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var address models.Address

	address.Id = uint(id)

	config.DB.Delete(&address)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
