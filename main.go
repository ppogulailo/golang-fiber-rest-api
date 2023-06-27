package main

import (
	"github.com/PogunGun/golang-fiber-rest-api/database"
	"github.com/PogunGun/golang-fiber-rest-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":3000")
}
