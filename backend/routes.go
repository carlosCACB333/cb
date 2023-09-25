package main

import (
	"cb/chatpdf"
	"cb/middleware"
	"cb/users"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "home_page.html", gin.H{})
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://carloscb.com",
			"https://www.carloscb.com",
		},
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "x-api-key"},
	}))

	{ // API V1
		{ // CONECTION ONLY BACK TO BACK
			v1Private := r.Group("/api/v1")
			v1Private.Use(middleware.ApiKeyMiddleware(os.Getenv("X_API_KEY")))
			{ //NOT REQUIRED AUTHENTICATION
				users.UserRoutes(v1Private.Group("/user"))
			}
			{ //REQUIRED AUTHENTICATION
				v1Private.Use(middleware.AuthMiddleware())
				chatpdf.ChatPdfRoutes(v1Private.Group("/chatpdf"))
				chatpdf.BootRoutesRoutes(v1Private.Group("/boot"))
				chatpdf.MessagesRoutes(v1Private.Group("/message"))
			}
		}

		{ // CONECTION BACK TO FRONT
			v1Public := r.Group("/api/v1/public")
			v1Public.Use(middleware.ApiKeyMiddleware(os.Getenv("X_API_KEY_PUBLIC")))
			{ //NOT REQUIRED AUTHENTICATION
				users.AuthRoutes(v1Public.Group("/auth"))
			}
			{ //REQUIRED AUTHENTICATION
				v1Public.Use(middleware.AuthMiddleware())
				v1Public.GET("/resource/:id", chatpdf.GetChatFile)
				v1Public.POST("/chatpdf", chatpdf.CreateChat)

			}
		}

	}

	return r
}
