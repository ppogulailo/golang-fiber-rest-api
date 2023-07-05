package controllers

import (
	"github.com/PogunGun/golang-fiber-rest-api/database"
	"github.com/PogunGun/golang-fiber-rest-api/models"
	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Role

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
