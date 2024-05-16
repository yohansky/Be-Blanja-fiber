package controller

import (
	"backend-gin/src/config"
	"backend-gin/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllPayments(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Payment{}, page))
}

func GetPaymentsByUserID(c *fiber.Ctx) error {
	id := c.Params("id")

	var payment []models.Payment
	if err := config.DB.Where("user_id = ?", id).Find(&payment).Error; err != nil {
		return c.JSON(fiber.Map{"Error": "User not found"})
	}

	return c.JSON(payment)
}

func CreatePayment(c *fiber.Ctx) error {
	var payment models.Payment

	if err := c.BodyParser(&payment); err != nil {
		return err
	}

	config.DB.Create(&payment)

	return c.JSON(payment)
}

func GetPayment(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var payment models.Payment

	payment.Id = uint(id)

	config.DB.Preload("User").Preload("Bank").Find(&payment)

	return c.JSON(payment)
}

func UpdatePayment(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var payment models.Payment

	payment.Id = uint(id)

	if err := c.BodyParser(&payment); err != nil {
		return err
	}

	config.DB.Model(&payment).Updates(payment)

	return c.JSON(payment)
}

func DeletePayment(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var payment models.Payment

	payment.Id = uint(id)

	config.DB.Delete(&payment)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
