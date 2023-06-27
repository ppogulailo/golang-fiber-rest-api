package routes

import (
	"github.com/PogunGun/golang-fiber-rest-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
	app.Get("/", controllers.Other)
}
