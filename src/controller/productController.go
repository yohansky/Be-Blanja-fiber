package controller

import (
	"backend-gin/src/config"
	"backend-gin/src/helper"
	"backend-gin/src/models"
	"backend-gin/src/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// func AllProducts(c *fiber.Ctx) error {
// 	page, _ := strconv.Atoi(c.Query("page", "1"))

// 	return c.JSON(models.Paginate(config.DB, &models.Product{}, page))
// }

func AllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

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

	file, err := c.FormFile("Image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Gagal mengunggah file: " + err.Error())
	}

	maxFileSize := int64(2 << 20)
	if err := helper.SizeUploadValidation(file.Size, maxFileSize); err != nil {
		return err
	}

	fileHeader, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal membuka file: " + err.Error())
	}
	defer fileHeader.Close()

	buffer := make([]byte, 512)
	if _, err := fileHeader.Read(buffer); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Gagal membaca file: " + err.Error())
	}

	validFileTypes := []string{"image/png", "image/jpeg", "image/jpg", "application/pdf"}
	if err := helper.TypeUploadValidation(buffer, validFileTypes); err != nil {
		return err
	}

	uploadResult, err := services.UploadCLoudinary(c, file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// if err := c.BodyParser(&product); err != nil {
	// 	return err
	// }

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// var product models.Product

	values := form.Value

	userID, err := strconv.ParseUint(values["UserId"][0], 10, 64)
	if err != nil {
		return err
	}

	price, err := strconv.ParseFloat(values["Price"][0], 64)
	if err != nil {
		return err
	}

	amount, err := strconv.ParseFloat(values["Amount"][0], 64)
	if err != nil {
		return err
	}

	product := models.Product{
		Name:        values["Name"][0],
		Price:       price,
		Color:       values["Color"][0],
		Size:        values["Size"][0],
		Amount:      amount,
		Condition:   values["Condition"][0],
		Description: values["Description"][0],
		Image:       uploadResult.SecureURL,
		UserId:      uint(userID),
	}

	config.DB.Create(&product)

	return c.JSON(fiber.Map{
		"Message": "Product created",
		"data":    product,
	})
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
