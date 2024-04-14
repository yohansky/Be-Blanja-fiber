package controller

import (
	"backend-gin/src/config"
	"backend-gin/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	config.DB.Preload("Permissions").Find(&roles)

	return c.JSON(roles)
}

func CreateRole(c *fiber.Ctx) error {
	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	//Many2many Relation
	//perhatikan string di dalam roleDto (harus sama dengan nama kolom di struct Role/ harus sama dengan nama struct Permissions)
	permissionsInterface, ok := roleDto["Permissions"].([]interface{})
	if !ok || permissionsInterface == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid or missing Permissions data",
		})
	}

	list := permissionsInterface

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDto["Name"].(string),
		Permissions: permissions,
	}

	config.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	// karena menggunakan Gorm.Model jadi tidak ada ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	// var role models.Role

	// role.ID = uint(id)

	config.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID format",
		})
	}
	var roleDto fiber.Map
	if err := c.BodyParser(&roleDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing request body",
		})
	}

	//Many2many Relation
	//perhatikan string di dalam roleDto (harus sama dengan nama kolom di struct Role/ harus sama dengan nama struct Permissions)
	permissionsInterface, ok := roleDto["Permissions"].([]interface{})
	if !ok || permissionsInterface == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid or missing Permissions data",
		})
	}

	list := permissionsInterface

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, err := strconv.Atoi(permissionId.(string))
		if err != nil {
			return err
		}

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	// var res interface{}

	// config.DB.Table("role_permissions").Where("role_id = ?", id).Delete(&res)
	config.DB.Exec("DELETE FROM role_permissions WHERE role_id = ?", id)

	role := models.Role{
		Id:          uint(id),
		Name:        roleDto["Name"].(string),
		Permissions: permissions,
	}

	config.DB.Model(&role).Updates(role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	config.DB.Exec("DELETE FROM role_permissions WHERE role_id = ?", id)

	role := models.Role{
		Id: uint(id),
	}

	// karena menggunakan Gorm.Model jadi tidak ada kolom ID, maknaya harus di deklarasikan dulu jika ingin dipanggil IDnya
	// var role models.Role

	// role.ID = uint(id)

	config.DB.Delete(&role)

	return c.JSON(fiber.Map{
		"Message": "Delete Complete",
	})
}
