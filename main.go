package main

import (
	"backend-gin/src/config"
	"backend-gin/src/helper"
	"backend-gin/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helper.Migrate()
	// defer config.DB.Close()
	app := fiber.New()
	routes.Router(app)
	app.Listen(":8080")
}
