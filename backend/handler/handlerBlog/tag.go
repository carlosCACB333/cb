package handlerBlog

import (
	"github.com/carlosCACB333/cb-back/dto"
	"github.com/carlosCACB333/cb-back/model"
	"github.com/carlosCACB333/cb-back/server"
	"github.com/carlosCACB333/cb-back/util"
	"github.com/gofiber/fiber/v2"
)

func CreateTag(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := s.DB()
		var tag model.Tag
		if err := c.BodyParser(&tag); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos inválidos", nil)
		}
		if err := util.ValidateFields(tag); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos inválidos", err)
		}
		tag.ID = util.NewID()
		if err := db.Create(&tag).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al crear la etiqueta", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    tag,
			Message: "Etiqueta creada correctamente",
		}))

	}
}

func UpdateTag(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := s.DB()
		var tagDTO dto.UpdateTag
		tagId := c.Params("id")
		if err := c.BodyParser(&tagDTO); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos inválidos", nil)
		}
		if err := util.ValidateFields(tagDTO); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos inválidos", err)
		}
		var tag model.Tag
		if err := db.Where("id = ?", tagId).First(&tag).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Etiqueta no encontrada", nil)
		}

		if err := db.Model(&tag).Updates(tagDTO).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al actualizar la etiqueta", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    tag,
			Message: "Etiqueta actualizada correctamente",
		}))

	}
}

func DeleteTag(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := s.DB()
		tagId := c.Params("id")
		var tag model.Tag
		if err := db.Where("id = ?", tagId).First(&tag).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Etiqueta no encontrada", nil)
		}
		if err := db.Delete(&tag).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al eliminar la etiqueta", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Message: "Etiqueta eliminada correctamente",
		}))

	}
}

func GetTagById(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := s.DB()
		tagId := c.Params("id")
		var tag model.Tag
		if err := db.Where("id = ?", tagId).First(&tag).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Etiqueta no encontrada", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data: tag,
		}))

	}
}

func GetTags(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := s.DB()
		var tags []model.Tag
		if err := db.Find(&tags).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al obtener las etiquetas", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data: tags,
		}))

	}
}
