package controllers

import (
	"github.com/PogunGun/golang-fiber-rest-api/database"
	"github.com/PogunGun/golang-fiber-rest-api/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllOrders(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}
