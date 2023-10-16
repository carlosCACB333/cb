package router

import (
	"github.com/carlosCACB333/cb-back/handler"
	"github.com/carlosCACB333/cb-back/middleware"

	"github.com/carlosCACB333/cb-back/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRouter(s *server.Server) {

	s.App().Static("/public", "./public")
	s.App().Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000,https://carloscb.com,https://www.carloscb.com",
		AllowHeaders: "*",
	}))

	s.App().Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{
			"Title": "carloscb",
		}, "layouts/main")
	})

	api := s.App().Group("/api")
	apiV1 := api.Group("/v1")

	setupPublicV1Router(s, apiV1)
	setupPrivatev1Router(s, apiV1)

}

func setupPrivatev1Router(s *server.Server, parent fiber.Router) {
	// CONECTION ONLY BACK TO BACK
	r := parent.Group("/", middleware.ApiKeyMiddleware(s.Config().ApiKey))
	{ //NOT REQUIRED AUTHENTICATION
		UserRouter(s, r.Group("/user"))
		BootRouterWithoutAuth(s,r.Group("/bootWithout"))
	}
	{ //REQUIRED AUTHENTICATION
		r.Use(middleware.AuthMiddleware(s))
		ChatPdfRouter(s, r.Group("/chatpdf"))
		BootRouter(s, r.Group("/boot"))
		MessagesRouter(s, r.Group("/message"))
		TagRouter(s, r.Group("/tag"))
		CategoryRouter(s, r.Group("/category"))
		PostRouter(s, r.Group("/post"))
		CommentRouter(s, r.Group("/comment"))
	}
}

func setupPublicV1Router(s *server.Server, parent fiber.Router) {
	// CONECTION FRONT TO BACK
	r := parent.Group("/public/", middleware.ApiKeyMiddleware(s.Config().ApiKeyPublic))
	{ //NOT REQUIRED AUTHENTICATION
		AuthRouter(s, r.Group("/auth"))
	}
	{ //REQUIRED AUTHENTICATION
		r.Use(middleware.AuthMiddleware(s))
		r.Get("/resource/:id", handler.GetChatFile(s))
		r.Post("/chatpdf", handler.CreateChat(s))
	}
}
