package routes

import (
	"github.com/PogunGun/golang-fiber-rest-api/controllers"
	"github.com/PogunGun/golang-fiber-rest-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthenticated)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.LogOut)
	// CRUD USER
	app.Get("/api/users", controllers.AllUser)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
}
