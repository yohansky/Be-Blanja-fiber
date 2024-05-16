package controller

import (
	"backend-gin/src/config"
	"backend-gin/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllBanks(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Bank{}, page))
	// var bank models.Bank

	// config.DB.Find(&bank)

	// return c.JSON(bank)
}

func CreateBank(c *fiber.Ctx) error {
	var order models.Bank

	if err := c.BodyParser(&order); err != nil {
		return err
	}

	config.DB.Create(&order)

	return c.JSON(order)
}

func DeleteBank(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var order models.Bank

	order.Id = uint(id)

	config.DB.Delete(&order)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
