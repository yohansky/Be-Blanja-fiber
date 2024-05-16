package routes

import (
	"backend-gin/src/controller"
	"backend-gin/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Post("/register", controller.Register)
	app.Post("/register/customer", controller.RegisterC)
	app.Post("/login", controller.Login)

	app.Use(middleware.IsAuth)

	app.Put("/user/info", controller.UpdateInfo)
	app.Put("/user/password", controller.UpdatePassword)

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

	app.Get("/products", controller.AllProducts)
	//all product by iduser
	app.Get("/products/user/:id", controller.GetProductsByUserID)
	app.Post("/products", controller.CreateProduct)
	app.Get("/product/:id", controller.GetProduct)
	app.Put("/product/:id", controller.UpdateProduct)
	app.Delete("/product/:id", controller.DeleteProduct)

	app.Get("/orders", controller.AllOrders)
	app.Get("/orders/user/:id", controller.GetOrdersByUserID)
	app.Get("/orders/user/:id/pending", controller.GetOrderByUserIDPending)
	app.Get("/order/:id", controller.GetOrder)
	app.Get("/order/user/:id", controller.GetOrderByUserID)
	app.Post("/orders", controller.CreateOrder)
	app.Put("/order/:id", controller.UpdateOrder)
	app.Delete("/order/:id", controller.DeleteOrder)

	app.Get("/address", controller.AllAddress)
	app.Get("/address/:id", controller.GetAddress)
	app.Get("/address/user/:id", controller.GetAddressByUserID)
	app.Post("/address", controller.CreateAddress)
	app.Put("/address/:id", controller.UpdateAddress)
	app.Delete("/address/:id", controller.DeleteAddress)

	app.Get("/delivery", controller.AllDelivery)
	app.Post("/delivery", controller.CreateDelivery)

	app.Get("/payments", controller.AllPayments)
	app.Get("/payment/:id", controller.GetPayment)
	app.Get("/payment/user/:id", controller.GetPaymentsByUserID)
	app.Post("/payments", controller.CreatePayment)
	app.Delete("/payment/:id", controller.DeletePayment)

	app.Get("/banks", controller.AllBanks)
	app.Post("/banks", controller.CreateBank)
	app.Delete("/bank/:id", controller.DeleteBank)

	app.Post("/upload", controller.Upload)
	app.Static("/uploads", "src/uploads")
}
