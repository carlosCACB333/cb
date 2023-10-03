package users

import (
	"cb/common"
	"cb/libs"
	"cb/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var userData SyncUserDTO
	if err := c.BodyParser(&userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Datos incorrectos",
		))
	}

	user := User{
		Model:     common.Model{ID: userData.ID},
		Email:     utils.NormalizeEmail(userData.Email),
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Photo:     userData.Photo,
		Phone:     userData.Phone,
		Status:    "active",
	}

	// find user
	var userFound User
	if err := libs.DBInit().Where("email = ?", user.Email).First(&userFound).Error; err == nil {
		// update user
		user.Status = "active"
		if err := libs.DBInit().Model(&userFound).Updates(user).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
				"error", "Error al actualizar usuario",
			))
		}
		return c.JSON(utils.ResponseMsg(
			"success", "Usuario actualizado correctamente",
		))
	}

	if err := libs.DBInit().Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al crear usuario",
		))
	}

	return c.JSON(utils.ResponseMsg(
		"success", "Usuario creado correctamente",
	))

}

func UpdateUser(c *fiber.Ctx) error {
	var userData SyncUserDTO
	if err := c.BodyParser(&userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Datos incorrectos",
		))
	}

	newUser := User{
		Model:     common.Model{ID: userData.ID},
		Email:     utils.NormalizeEmail(userData.Email),
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Photo:     userData.Photo,
		Phone:     userData.Phone,
		Status:    "active",
	}

	// get user
	var user User
	if err := libs.DBInit().Where("id = ?", newUser.ID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.ResponseMsg(
			"error", "Usuario no encontrado",
		))
	}

	// update user
	if err := libs.DBInit().Model(&user).Updates(newUser).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al actualizar usuario",
		))

	}

	return c.JSON(utils.ResponseMsg(
		"success", "Usuario actualizado correctamente",
	))

}

func DeleUser(c *fiber.Ctx) error {
	id := c.Params("id")

	// get user
	var user User
	if err := libs.DBInit().Where("id = ?", id).First(&user).Error; err != nil {

		return c.Status(fiber.StatusNotFound).JSON(utils.ResponseMsg(
			"error", "Usuario no encontrado",
		))

	}

	// update status
	if err := libs.DBInit().Model(&user).Updates(User{Status: "deleted"}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al eliminar usuario",
		))
	}
	return c.JSON(utils.ResponseMsg(
		"success", "Usuario eliminado correctamente",
	))
}
