package router

import (
	blogHandler "github.com/carlosCACB333/cb-back/handler/handlerBlog"
	"github.com/carlosCACB333/cb-back/server"
	"github.com/gofiber/fiber/v2"
)

func TagRouter(s *server.Server, r fiber.Router) {
	r.Get("/", blogHandler.GetTags(s))
	r.Get("/:id", blogHandler.GetTagById(s))
	r.Post("/", blogHandler.CreateTag(s))
	r.Put("/:id", blogHandler.UpdateTag(s))
	r.Delete("/:id", blogHandler.DeleteTag(s))
}

func CategoryRouter(s *server.Server, r fiber.Router) {
	r.Get("/", blogHandler.GetCategories(s))
	r.Get("/:id", blogHandler.GetCategoryById(s))
	r.Post("/", blogHandler.CreateCategory(s))
	r.Put("/:id", blogHandler.UpdateCategory(s))
	r.Delete("/:id", blogHandler.DeleteCategory(s))
}

func PostRouter(s *server.Server, r fiber.Router) {
	r.Post("/", blogHandler.CreatePost(s))
	r.Get("/", blogHandler.GetPosts(s))
	r.Get("/:id", blogHandler.GetPostById(s))
	r.Put("/:id", blogHandler.UpdatePost(s))
	r.Delete("/:id", blogHandler.DeletePost(s))
}

func CommentRouter(s *server.Server, r fiber.Router) {
	r.Post("/", blogHandler.CreateComment(s))
	r.Get("/", blogHandler.GetComments(s))
	r.Get("/:id", blogHandler.GetCommentById(s))
	r.Put("/:id", blogHandler.UpdateComment(s))
	r.Delete("/:id", blogHandler.DeleteComment(s))
}
