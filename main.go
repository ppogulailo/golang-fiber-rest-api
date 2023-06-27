package main

import (
	"github.com/PogunGun/golang-fiber-rest-api/database"
	"github.com/PogunGun/golang-fiber-rest-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)
	app.Listen(":3000")
}
