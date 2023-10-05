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
		return utils.NewError(fiber.StatusBadRequest, "Datos incorrectos", nil)
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
			return utils.NewError(fiber.StatusBadRequest, "Error al actualizar usuario", nil)
		}
		return c.JSON(utils.NewBody(utils.Body{
			Message: "Usuario actualizado correctamente",
		}))

	}

	if err := libs.DBInit().Create(&user).Error; err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Error al crear usuario", nil)
	}

	return c.JSON(utils.NewBody(utils.Body{
		Message: "Usuario creado correctamente",
	}))

}

func UpdateUser(c *fiber.Ctx) error {
	var userData SyncUserDTO
	if err := c.BodyParser(&userData); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Datos incorrectos", nil)

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
		return utils.NewError(fiber.StatusNotFound, "Usuario no encontrado", nil)
	}

	// update user
	if err := libs.DBInit().Model(&user).Updates(newUser).Error; err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Error al actualizar usuario", nil)

	}

	return c.JSON(utils.NewBody(utils.Body{
		Message: "Usuario actualizado correctamente",
	}))

}

func DeleUser(c *fiber.Ctx) error {
	id := c.Params("id")

	// get user
	var user User
	if err := libs.DBInit().Where("id = ?", id).First(&user).Error; err != nil {
		return utils.NewError(fiber.StatusNotFound, "Usuario no encontrado", nil)

	}

	// update status
	if err := libs.DBInit().Model(&user).Updates(User{Status: "deleted"}).Error; err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Error al eliminar usuario", nil)
	}

	return c.JSON(utils.NewBody(utils.Body{
		Message: "Usuario eliminado correctamente",
	}))

}
