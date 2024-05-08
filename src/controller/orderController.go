package controller

import (
	"backend-gin/src/config"
	"backend-gin/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllOrders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Order{}, page))
}

func GetOrdersByUserID(c *fiber.Ctx) error {
	id := c.Params("id")

	var order []models.Order
	if err := config.DB.Where("user_id = ?", id).Preload("Product").Find(&order).Error; err != nil {
		return c.JSON(fiber.Map{"Error": "User not found"})
	}

	return c.JSON(order)
}

func GetOrderByUserID(c *fiber.Ctx) error {
	id := c.Params("id")

	var order []models.Order
	if err := config.DB.Where("user_id = ?", id).First(&order).Error; err != nil {
		return c.JSON(fiber.Map{"Error": "User not found"})
	}

	return c.JSON(order)
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return err
	}

	if err := order.BeforeSave(config.DB); err != nil {
		return err
	}

	config.DB.Create(&order)

	return c.JSON(order)
}

func GetOrderPending(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var order models.Order

	order.Id = uint(id)

	if err := order.BeforeSave(config.DB); err != nil {
		return err
	}

	config.DB.Preload("Product").Find(&order)

	return c.JSON(order)
}

func GetOrder(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var order models.Order

	order.Id = uint(id)

	config.DB.Preload("Product").Find(&order)

	return c.JSON(order)
}

func UpdateOrder(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var order models.Order

	order.Id = uint(id)

	if err := c.BodyParser(&order); err != nil {
		return err
	}

	config.DB.Model(&order).Updates(order)

	return c.JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var order models.Order

	order.Id = uint(id)

	config.DB.Delete(&order)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
