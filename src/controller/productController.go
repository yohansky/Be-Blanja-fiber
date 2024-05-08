package controller

import (
	"backend-gin/src/config"
	"backend-gin/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// func AllProducts(c *fiber.Ctx) error {
// 	page, _ := strconv.Atoi(c.Query("page", "1"))

// 	return c.JSON(models.Paginate(config.DB, &models.Product{}, page))
// }

func AllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	// Preload "User" relation
	products := []models.Product{}
	result := config.DB.Preload("User").Find(&products)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch products",
		})
	}

	return c.JSON(models.PaginateProducts(config.DB, &models.Product{}, page))
}

func GetProductsByUserID(c *fiber.Ctx) error {
	id := c.Params("id")

	var product []models.Product
	if err := config.DB.Where("user_id = ?", id).Find(&product).Error; err != nil {
		return c.JSON(fiber.Map{"Error": "User not found"})
	}

	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	config.DB.Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var product models.Product

	product.Id = uint(id)

	config.DB.Preload("User").Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var product models.Product

	product.Id = uint(id)

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	config.DB.Model(&product).Updates(product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	var product models.Product

	product.Id = uint(id)

	config.DB.Delete(&product)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
