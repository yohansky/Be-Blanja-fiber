package controller

import (
	"backend-gin/src/config"
	"backend-gin/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllDelivery(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(config.DB, &models.Delivery{}, page))
}

func CreateDelivery(c *fiber.Ctx) error {
	var delivery models.Delivery

	if err := c.BodyParser(&delivery); err != nil {
		return err
	}

	// if err := delivery.BeforeSave(config.DB); err != nil {
	// 	return err
	// }

	config.DB.Create(&delivery)

	return c.JSON(delivery)
}
