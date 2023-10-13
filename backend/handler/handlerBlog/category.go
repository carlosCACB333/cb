package handlerBlog

import (
	"github.com/carlosCACB333/cb-back/dto"
	"github.com/carlosCACB333/cb-back/model"
	"github.com/carlosCACB333/cb-back/server"
	"github.com/carlosCACB333/cb-back/util"
	"github.com/gofiber/fiber/v2"
)

func CreateCategory(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var category model.Category
		if err := c.BodyParser(&category); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos inválidos", nil)
		}
		category.ID = util.NewID()
		category.Slug = util.Slug(category.Name)
		if err := util.ValidateFields(category); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos incompletos", err)
		}
		if err := s.DB().Create(&category).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al crear la categoría", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    category,
			Message: "Categoría creada correctamente",
		}))

	}

}

func UpdateCategory(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := s.DB()
		var categoryDto dto.UpdateCategory
		categoryId := c.Params("id")
		if err := c.BodyParser(&categoryDto); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos inválidos", nil)
		}
		categoryDto.Slug = util.Slug(categoryDto.Name)
		if err := util.ValidateFields(categoryDto); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos incompletos", err)
		}
		var category model.Category
		if err := db.Where("id = ?", categoryId).First(&category).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Categoría no encontrada", nil)
		}

		if err := db.Model(&category).Updates(categoryDto).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al actualizar la categoría", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    category,
			Message: "Categoría actualizada correctamente",
		}))

	}

}

func DeleteCategory(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := s.DB()
		categoryId := c.Params("id")
		var category model.Category
		if err := db.Where("id = ?", categoryId).First(&category).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Categoría no encontrada", nil)
		}

		if err := db.Delete(&category).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al eliminar la categoría", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    category,
			Message: "Categoría eliminada correctamente",
		}))

	}

}

func GetCategoryById(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := s.DB()
		categoryId := c.Params("id")
		var category model.Category
		if err := db.Where("id = ?", categoryId).First(&category).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Categoría no encontrada", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    category,
			Message: "Categoría encontrada",
		}))

	}

}

func GetCategories(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := s.DB()
		var categories []model.Category
		if err := db.Find(&categories).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al obtener las categorías", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    categories,
			Message: "Categorías encontradas",
		}))

	}

}
