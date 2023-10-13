package handlerBlog

import (
	"fmt"

	"github.com/carlosCACB333/cb-back/dto"
	"github.com/carlosCACB333/cb-back/model"
	"github.com/carlosCACB333/cb-back/server"
	"github.com/carlosCACB333/cb-back/util"
	"github.com/gofiber/fiber/v2"
)

func CreateComment(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)
		var comment model.Comment
		if err := c.BodyParser(&comment); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos inválidos", nil)
		}
		comment.ID = util.NewID()
		comment.AuthorId = user.ID
		if err := util.ValidateFields(comment); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos incompletos", err)
		}

		if err := s.DB().Create(&comment).Error; err != nil {
			fmt.Println(err)
			return util.NewError(fiber.StatusBadRequest, "Error al crear el comentario", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    comment,
			Message: "Comentario creado correctamente",
		}))

	}

}

func GetComments(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var comments []model.Comment
		if err := s.DB().Find(&comments).Order("created_at desc").Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al obtener los comentarios", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    comments,
			Message: "Comentarios obtenidos correctamente",
		}))
	}
}

func GetCommentById(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var comment model.Comment
		if err := s.DB().Where("id = ?", id).First(&comment).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al obtener el comentario", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    comment,
			Message: "Comentario obtenido correctamente",
		}))
	}
}

func UpdateComment(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		user := c.Locals("user").(*model.User)
		var commentDto dto.UpdateComment

		if err := c.BodyParser(&commentDto); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos inválidos", nil)
		}
		var comment model.Comment
		if err := s.DB().Where("id = ?", id).First(&comment).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al obtener el comentario", nil)
		}

		if comment.AuthorId != user.ID {
			return util.NewError(fiber.StatusUnauthorized, "No tienes permisos para actualizar este comentario", nil)
		}

		if s.DB().Model(&comment).Updates(commentDto).Error != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al actualizar el comentario", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    comment,
			Message: "Comentario actualizado correctamente",
		}))
	}
}

func DeleteComment(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		user := c.Locals("user").(*model.User)
		var comment model.Comment
		if err := s.DB().Where("id = ?", id).First(&comment).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al obtener el comentario", nil)
		}

		if comment.AuthorId != user.ID {
			return util.NewError(fiber.StatusUnauthorized, "No tienes permisos para eliminar este comentario", nil)
		}

		if s.DB().Delete(&comment).Error != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al eliminar el comentario", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    comment,
			Message: "Comentario eliminado correctamente",
		}))
	}
}
