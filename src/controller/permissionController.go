package controller

import (
	"backend-gin/src/config"
	"backend-gin/src/models"

	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	config.DB.Find(&permissions)

	return c.JSON(permissions)
}
