package routes

import (
	"backend-gin/src/controller"
	"backend-gin/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)

	app.Use(middleware.IsAuth)

	app.Get("/user", controller.User)
	app.Post("/logout", controller.Logout)

	app.Get("/users", controller.AllUsers)
	app.Post("/users", controller.CreateUser)
	app.Get("/user/:id", controller.GetUser)
	app.Put("/user/:id", controller.UpdateUser)
	app.Delete("/user/:id", controller.DeleteUser)

	app.Get("/roles", controller.AllRoles)
	app.Post("/roles", controller.CreateRole)
	app.Get("/role/:id", controller.GetRole)
	app.Put("/role/:id", controller.UpdateRole)
	app.Delete("/role/:id", controller.DeleteRole)

	app.Get("/permissions", controller.AllPermissions)
}
