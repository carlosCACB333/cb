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
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Datos incorrectos",
		))
	}
	// validate fields
	errors := utils.ValidateFields(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response(
			"error", "Campos invalido",
			errors,
			nil,
		))
	}

	user.ID = uuid.New().String()
	user.Email = utils.NormalizeEmail(user.Email)
	user.Password = libs.HashPassword(user.Password)
	user.Status = "active"

	// create user
	if err := libs.DBInit().Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al crear usuario",
		))
	}

	return c.JSON(utils.Response(
		"success", "Usuario creado correctamente",
		user,
		nil,
	))

}

func AuthLogin(c *fiber.Ctx) error {
	var login LoginDTO
	var user User

	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Datos incorrectos",
		))
	}

	// validate fields
	errors := utils.ValidateFields(login)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response(
			"error", "Campos invalidos",
			errors,
			nil,
		))
	}

	login.Email = utils.NormalizeEmail(login.Email)

	// check if user exists
	if err := libs.DBInit().Where("email = ?", login.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Usuario no encontrado",
		))
	}

	// check if password is correct
	if !libs.CheckPassword(login.Password, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Contraseña incorrecta",
		))
	}

	// generate token
	token, tkerr := libs.GenerateToken(user.ID)
	if tkerr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al generar token",
		))

	}

	return c.JSON(utils.Response(
		"success", "Login successful",
		fiber.Map{"token": token, "user": user},
		nil,
	))

}

func ChangePassword(c *fiber.Ctx) error {
	clearkUser := c.Locals("user").(*User)
	var reset ChangePasswordDTO
	if err := c.BodyParser(&reset); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Datos incorrectos",
		))
	}

	// validate fields
	if err := utils.ValidateFields(reset); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response(
			"error", "Campos invalidos",
			err,
			nil,
		))
	}
	if reset.NewPassword != reset.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Las contraseñas no coinciden",
		))
	}
	var user User
	if err := libs.DBInit().Where("id = ?", clearkUser.ID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Usuario no encontrado",
		))
	}

	// check if old password is correct
	if !libs.CheckPassword(reset.OldPassword, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Contraseña antigua incorrecta",
		))
	}

	// update password
	user.Password = libs.HashPassword(reset.NewPassword)
	if err := libs.DBInit().Save(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseMsg(
			"error", "Error al actualizar contraseña",
		))
	}

	return c.JSON(utils.Response(
		"success", "Contraseña actualizada correctamente",
		user,
		nil,
	))

}
