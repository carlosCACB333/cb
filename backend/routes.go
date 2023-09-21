package main

import (
	"cb/chatpdf"
	"cb/middleware"
	"cb/users"

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
		AllowOrigins: []string{"http://localhost:3000", "https://carloscb.com"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "x-api-key"},
	}))

	r.Use(middleware.ApiKeyMiddleware())
	{
		v1 := r.Group("/api/v1")
		{
			//PUBLIC ROUTES
			users.AuthRoutes(v1.Group("/auth"))
			users.UserRoutes(v1.Group("/user"))
			v1.POST("/test/:id", chatpdf.Test)
		}

		{
			//PRIVATE ROUTES
			v1.Use(middleware.AuthMiddleware())
			chatpdf.ChatPdfRoutes(v1.Group("/chatpdf"))
			chatpdf.BootRoutesRoutes(v1.Group("/boot"))
			chatpdf.MessagesRoutes(v1.Group("/message"))
		}
	}

	return r
}
