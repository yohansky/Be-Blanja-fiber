package routes

import (
	"backend-gin/src/controller"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/", controller.Hello)
	app.Get("/other", controller.Other)
}
