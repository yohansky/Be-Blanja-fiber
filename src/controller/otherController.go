package controller

import "github.com/gofiber/fiber/v2"

func Other(c *fiber.Ctx) error {
	return c.SendString("Other COntroller")
}
