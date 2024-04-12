package routes

import (
	"backend-gin/src/controller"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)
}
