package controllers

import (
	"github.com/PogunGun/golang-fiber-rest-api/database"
	"github.com/PogunGun/golang-fiber-rest-api/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(roles)
}

type RoleCreateDTO struct {
	name        string
	permissions []string
}

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}
	list := roleDto["permissions"].([]interface{})
	permission := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permission[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:       roleDto["name"].(string),
		Permission: permission,
	}
	database.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	role := models.Role{
		Id: uint(id),
	}
	database.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var roleDto fiber.Map
	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}
	list := roleDto["permissions"].([]interface{})
	permission := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permission[i] = models.Permission{
			Id: uint(id),
		}
	}
	var result interface{}
	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)
	role := models.Role{
		Id:         uint(id),
		Name:       roleDto["name"].(string),
		Permission: permission,
	}
	database.DB.Model(&role).Updates(role)
	return c.JSON(role)
}
func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	role := models.Role{
		Id: uint(id),
	}
	if err := c.BodyParser(&role); err != nil {
		return err
	}
	database.DB.Delete(&role)
	return nil
}
