package controller

import (
	"backend-gin/src/config"
	"backend-gin/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllCategories(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Category{}, page))
}

func CreateCategory(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	name := form.Value["Name"][0]
	files := form.File["Image"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No image file uploaded",
		})
	}

	file := files[0]

	if err := c.SaveFile(file, "src/uploads/"+file.Filename); err != nil {
		return err
	}

	category := models.Category{
		Name:  name,
		Image: "http://localhost:8080/uploads/" + file.Filename,
	}

	config.DB.Create(&category)

	return c.JSON(fiber.Map{
		"message": "Category created successfully",
		"data":    category,
	})
}
