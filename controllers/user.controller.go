package controllers

import (
	"github.com/PogunGun/golang-fiber-rest-api/database"
	"github.com/PogunGun/golang-fiber-rest-api/models"
	"github.com/gofiber/fiber/v2"
	"math"
	"strconv"
)

func AllUser(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	limit := 5
	offset := (page - 1) * limit
	var toltal int64

	var users []models.User

	database.DB.Preload("Role").Offset(offset).Limit(limit).Find(&users)

	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"total":     toltal,
			"page":      page,
			"last_page": math.Ceil(float64(int(toltal) / limit)),
		},
	})
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.SetPassword("1234")
	database.DB.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		Id: uint(id),
	}
	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		Id: uint(id),
	}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	database.DB.Model(&user).Updates(user)
	return c.JSON(user)
}
func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		Id: uint(id),
	}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	database.DB.Delete(&user)
	return nil
}
