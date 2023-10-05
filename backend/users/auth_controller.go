package users

import (
	"cb/libs"
	"cb/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AuthRegister(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Datos incorrectos", nil)
	}
	// validate fields
	errors := utils.ValidateFields(user)
	if errors != nil {
		return utils.NewError(fiber.StatusBadRequest, "Campos invalidos", errors)
	}

	user.ID = uuid.New().String()
	user.Email = utils.NormalizeEmail(user.Email)
	user.Password = libs.HashPassword(user.Password)
	user.Status = "active"

	// create user
	if resp := libs.DBInit().Create(&user); resp.Error != nil {
		return utils.NewError(fiber.StatusBadRequest, "Error al crear usuario", nil)
	}

	return c.JSON(utils.NewBody(utils.Body{
		Message: "Usuario creado correctamente",
		Data:    user,
	}))

}

func AuthLogin(c *fiber.Ctx) error {
	var login LoginDTO
	var user User

	if err := c.BodyParser(&login); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Datos incorrectos", nil)
	}

	// validate fields
	errors := utils.ValidateFields(login)
	if errors != nil {
		return utils.NewError(fiber.StatusBadRequest, "Campos invalidos", errors)
	}

	login.Email = utils.NormalizeEmail(login.Email)

	// check if user exists
	if err := libs.DBInit().Where("email = ?", login.Email).First(&user).Error; err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Usuario no encontrado", nil)
	}

	// check if password is correct
	if !libs.CheckPassword(login.Password, user.Password) {
		return utils.NewError(fiber.StatusBadRequest, "Contraseña incorrecta", nil)
	}

	// generate token
	token, tkerr := libs.GenerateToken(user.ID)
	if tkerr != nil {
		return utils.NewError(fiber.StatusBadRequest, "Error al generar token", nil)
	}
	return c.JSON(utils.NewBody(utils.Body{
		Message: "Login successful",
		Data:    fiber.Map{"token": token, "user": user},
	}))

}

func ChangePassword(c *fiber.Ctx) error {
	clearkUser := c.Locals("user").(*User)
	var reset ChangePasswordDTO
	if err := c.BodyParser(&reset); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Datos incorrectos", nil)
	}

	// validate fields
	if err := utils.ValidateFields(reset); err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Campos invalidos", err)
	}
	if reset.NewPassword != reset.ConfirmPassword {
		return utils.NewError(fiber.StatusBadRequest, "Las contraseñas no coinciden", nil)
	}
	var user User
	if err := libs.DBInit().Where("id = ?", clearkUser.ID).First(&user).Error; err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Usuario no encontrado", nil)
	}

	// check if old password is correct
	if !libs.CheckPassword(reset.OldPassword, user.Password) {
		return utils.NewError(fiber.StatusBadRequest, "Contraseña antigua incorrecta", nil)
	}

	// update password
	user.Password = libs.HashPassword(reset.NewPassword)
	if err := libs.DBInit().Save(&user).Error; err != nil {
		return utils.NewError(fiber.StatusBadRequest, "Error al actualizar contraseña", nil)
	}

	return c.JSON(utils.NewBody(utils.Body{
		Message: "Contraseña actualizada correctamente",
		Data:    user,
	}))

}
