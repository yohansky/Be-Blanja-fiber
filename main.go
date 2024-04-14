package main

import (
	"backend-gin/src/config"
	"backend-gin/src/helper"
	"backend-gin/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helper.Migrate()
	// defer config.DB.Close()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	routes.Router(app)
	app.Listen(":8080")
}
