package handler

import (
	"fmt"

	"github.com/carlosCACB333/cb-back/dto"
	"github.com/carlosCACB333/cb-back/lib"
	"github.com/carlosCACB333/cb-back/model"
	"github.com/carlosCACB333/cb-back/server"
	"github.com/carlosCACB333/cb-back/util"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	server.Server
}

func NewAuthHandler(s *server.Server) *AuthHandler {
	return &AuthHandler{*s}
}

func (h *AuthHandler) AuthRegister(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return util.NewError(fiber.StatusBadRequest, "Datos incorrectos", nil)
	}
	// validate fields
	errors := util.ValidateFields(user)
	if errors != nil {
		return util.NewError(fiber.StatusBadRequest, "Campos invalidos", errors)
	}

	user.ID = util.NewID()
	user.Email = util.NormalizeEmail(user.Email)
	user.Password = lib.HashPassword(user.Password)
	user.Status = "active"

	// create user
	if resp := h.DB().Create(&user); resp.Error != nil {
		return util.NewError(fiber.StatusBadRequest, "Error al crear usuario", nil)
	}

	return c.JSON(util.NewBody(util.Body{
		Message: "Usuario creado correctamente",
		Data:    user,
	}))

}

func (h *AuthHandler) AuthLogin(c *fiber.Ctx) error {
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
	if err := h.DB().Where("email = ?", login.Email).First(&user).Error; err != nil {
		fmt.Println(err)
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

func (h *AuthHandler) ChangePassword(c *fiber.Ctx) error {

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
	if err := h.DB().Where("id = ?", clearkUser.ID).First(&user).Error; err != nil {
		return util.NewError(fiber.StatusBadRequest, "Usuario no encontrado", nil)
	}

	// check if old password is correct
	if !lib.CheckPassword(reset.OldPassword, user.Password) {
		return util.NewError(fiber.StatusBadRequest, "Contraseña antigua incorrecta", nil)
	}

	// update password
	user.Password = lib.HashPassword(reset.NewPassword)
	if err := h.DB().Save(&user).Error; err != nil {
		return util.NewError(fiber.StatusBadRequest, "Error al actualizar contraseña", nil)
	}

	return c.JSON(util.NewBody(util.Body{
		Message: "Contraseña actualizada correctamente",
		Data:    user,
	}))

}
