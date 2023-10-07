package handler

import (
	"cb/dto"
	"cb/model"
	"cb/server"
	"cb/util"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userData dto.SyncUserDTO
		if err := c.BodyParser(&userData); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos incorrectos", nil)
		}

		user := model.User{
			Model:     model.Model{ID: userData.ID},
			Email:     util.NormalizeEmail(userData.Email),
			FirstName: userData.FirstName,
			LastName:  userData.LastName,
			Photo:     userData.Photo,
			Phone:     userData.Phone,
			Status:    "active",
		}

		// find user
		var userFound model.User
		if err := s.DB().Where("email = ?", user.Email).First(&userFound).Error; err == nil {
			// update user
			user.Status = "active"
			if err := s.DB().Model(&userFound).Updates(user).Error; err != nil {
				return util.NewError(fiber.StatusBadRequest, "Error al actualizar usuario", nil)
			}
			return c.JSON(util.NewBody(util.Body{
				Message: "Usuario actualizado correctamente",
			}))

		}

		if err := s.DB().Create(&user).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al crear usuario", nil)
		}

		return c.JSON(util.NewBody(util.Body{
			Message: "Usuario creado correctamente",
		}))

	}
}

func UpdateUser(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userData dto.SyncUserDTO
		if err := c.BodyParser(&userData); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos incorrectos", nil)

		}
		newUser := model.User{
			Model:     model.Model{ID: userData.ID},
			Email:     util.NormalizeEmail(userData.Email),
			FirstName: userData.FirstName,
			LastName:  userData.LastName,
			Photo:     userData.Photo,
			Phone:     userData.Phone,
			Status:    "active",
		}

		// get user
		var user model.User
		if err := s.DB().Where("id = ?", newUser.ID).First(&user).Error; err != nil {
			return util.NewError(fiber.StatusNotFound, "Usuario no encontrado", nil)
		}

		// update user
		if err := s.DB().Model(&user).Updates(newUser).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al actualizar usuario", nil)

		}

		return c.JSON(util.NewBody(util.Body{
			Message: "Usuario actualizado correctamente",
		}))

	}
}

func DeleUser(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		// get user
		var user model.User
		if err := s.DB().Where("id = ?", id).First(&user).Error; err != nil {
			return util.NewError(fiber.StatusNotFound, "Usuario no encontrado", nil)

		}

		// update status
		if err := s.DB().Model(&user).Updates(model.User{Status: "deleted"}).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al eliminar usuario", nil)
		}

		return c.JSON(util.NewBody(util.Body{
			Message: "Usuario eliminado correctamente",
		}))

	}
}
