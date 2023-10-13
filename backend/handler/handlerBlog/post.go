package handlerBlog

import (
	"github.com/carlosCACB333/cb-back/dto"
	"github.com/carlosCACB333/cb-back/event"
	"github.com/carlosCACB333/cb-back/model"
	"github.com/carlosCACB333/cb-back/server"
	"github.com/carlosCACB333/cb-back/util"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)
		var post model.Post
		if err := c.BodyParser(&post); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos inválidos", nil)
		}
		post.ID = util.NewID()
		post.AuthorId = user.ID
		post.Slug = util.Slug(post.Title)
		if err := util.ValidateFields(post); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos incompletos", err)
		}

		for i := range post.Tags {
			tag := *post.Tags[i]
			if tag.ID == "" {
				tag.ID = util.NewID()
				if err := util.ValidateFields(tag); err != nil {
					return util.NewError(fiber.StatusBadRequest, "Datos incompletos",
						map[string]any{"tags": err},
					)
				}
				post.Tags[i] = &tag
			}
		}

		if err := s.DB().Create(&post).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al crear el post", nil)
		}

		event.NewPostEvent(s.Nats()).PublishCreated(post)

		return c.JSON(util.NewBody(util.Body{
			Data:    post,
			Message: "Post creado correctamente",
		}))

	}

}

func GetPosts(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var posts []model.Post
		if err := s.DB().
			Preload("Category").
			Preload("Tags").
			Omit("Author").
			Order("created_at desc").
			Find(&posts).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al obtener los posts", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    posts,
			Message: "Posts obtenidos correctamente",
		}))

	}

}

func GetPostById(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		postId := c.Params("id")
		var post model.Post
		if err := s.DB().
			Where("id = ?", postId).
			Preload("Category").
			Preload("Tags").
			First(&post).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Post no encontrado", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    post,
			Message: "Post obtenido correctamente",
		}))

	}

}

func UpdatePost(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)
		db := s.DB()
		postId := c.Params("id")
		var postDTO dto.UpdatePost
		if err := c.BodyParser(&postDTO); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos inválidos", nil)
		}
		var post model.Post
		if err := db.Where("id = ?", postId).First(&post).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Post no encontrado", nil)
		}

		if post.AuthorId != user.ID {
			return util.NewError(fiber.StatusUnauthorized, "No tienes permisos para actualizar este post", nil)
		}

		postDTO.Slug = util.Slug(postDTO.Title)
		if err := util.ValidateFields(postDTO); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Datos inválidos", err)
		}

		for i := range postDTO.Tags {
			tag := *postDTO.Tags[i]
			if tag.ID == "" {
				tag.ID = util.NewID()
				if err := util.ValidateFields(tag); err != nil {
					return util.NewError(fiber.StatusBadRequest, "Datos incompletos",
						map[string]any{"tags": err},
					)
				}
				postDTO.Tags[i] = &tag
			}
		}

		if err := db.Model(&post).Association("Tags").Replace(postDTO.Tags); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al actualizar las etiquetas", nil)
		}

		if err := db.Model(&post).Updates(postDTO).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al actualizar el post", nil)
		}
		return c.JSON(util.NewBody(util.Body{
			Data:    post,
			Message: "Post actualizado correctamente",
		}))

	}
}

func DeletePost(s *server.Server) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)
		db := s.DB()
		postId := c.Params("id")
		var post model.Post

		if err := db.Where("id = ?", postId).First(&post).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Post no encontrado", nil)
		}

		if post.AuthorId != user.ID {
			return util.NewError(fiber.StatusUnauthorized, "No tienes permisos para eliminar este post", nil)
		}

		if err := db.Model(&post).Association("Tags").Clear(); err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al eliminar las etiquetas", nil)
		}

		if err := db.Delete(&post).Error; err != nil {
			return util.NewError(fiber.StatusBadRequest, "Error al eliminar el post", nil)
		}

		return c.JSON(util.NewBody(util.Body{
			Message: "Post eliminado correctamente",
		}))

	}

}
