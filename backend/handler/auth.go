package handler

import (
	"cb/dto"
	"cb/lib"
	"cb/model"
	"cb/server"
	"cb/util"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AuthRegister(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user model.User
		if err := c.BodyParser(&user); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos incorrectos", nil)
		}
		// validate fields
		errors := util.ValidateFields(user)
		if errors != nil {
			return util.NewError(fiber.StatusBadRequest, "Campos invalidos", errors)
		}

		user.ID = uuid.New().String()
		user.Email = util.NormalizeEmail(user.Email)
		user.Password = lib.HashPassword(user.Password)
		user.Status = "active"

		// create user
		if resp := s.DB().Create(&user); resp.Error != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al crear usuario", nil)
		}

		return c.JSON(util.NewBody(util.Body{
			Message: "Usuario creado correctamente",
			Data:    user,
		}))

	}
}

func AuthLogin(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var login dto.LoginDTO
		var user model.User

		if err := c.BodyParser(&login); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos incorrectos", nil)
		}

		// validate fields
		errors := util.ValidateFields(login)
		if errors != nil {
			return util.NewError(fiber.StatusBadRequest, "Campos invalidos", errors)
		}

		login.Email = util.NormalizeEmail(login.Email)

		// check if user exists
		if err := s.DB().Where("email = ?", login.Email).First(&user).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Usuario no encontrado", nil)
		}

		// check if password is correct
		if !lib.CheckPassword(login.Password, user.Password) {
			return util.NewError(fiber.StatusBadRequest, "Contraseña incorrecta", nil)
		}

		// generate token
		token, tkerr := lib.GenerateToken(user.ID)
		if tkerr != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al generar token", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Message: "Login successful",
			Data:    fiber.Map{"token": token, "user": user},
		}))

	}
}

func ChangePassword(s *server.Server) fiber.Handler {

	return func(c *fiber.Ctx) error {
		clearkUser := c.Locals("user").(*model.User)
		var reset dto.ChangePasswordDTO
		if err := c.BodyParser(&reset); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos incorrectos", nil)
		}

		// validate fields
		if err := util.ValidateFields(reset); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Campos invalidos", err)
		}
		if reset.NewPassword != reset.ConfirmPassword {
			return util.NewError(fiber.StatusBadRequest, "Las contraseñas no coinciden", nil)
		}
		var user model.User
		if err := s.DB().Where("id = ?", clearkUser.ID).First(&user).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Usuario no encontrado", nil)
		}

		// check if old password is correct
		if !lib.CheckPassword(reset.OldPassword, user.Password) {
			return util.NewError(fiber.StatusBadRequest, "Contraseña antigua incorrecta", nil)
		}

		// update password
		user.Password = lib.HashPassword(reset.NewPassword)
		if err := s.DB().Save(&user).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al actualizar contraseña", nil)
		}

		return c.JSON(util.NewBody(util.Body{
			Message: "Contraseña actualizada correctamente",
			Data:    user,
		}))

	}
}
