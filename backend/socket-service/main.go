package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	C "github.com/carlosCACB333/cb-back/const"
	"github.com/carlosCACB333/cb-back/event"
	"github.com/carlosCACB333/cb-back/middleware"
	"github.com/carlosCACB333/cb-back/model"
	"github.com/carlosCACB333/cb-back/server"
	"github.com/carlosCACB333/cb-back/util"
	ws "github.com/carlosCACB333/cb-back/websocket"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
)

func main() {

	s := server.New(
		context.Background(),
		server.Config{
			Config: fiber.Config{
				CaseSensitive: true,
				AppName:       "Websocket Service",
				// ErrorHandler:  middleware.ErrorHandler,
				IdleTimeout:  time.Second * 20,
				ReadTimeout:  time.Second * 10,
				WriteTimeout: time.Second * 10,
			},
			Port:         ":" + os.Getenv("PORT"),
			Dialector:    postgres.Open(os.Getenv("DB_URL")),
			ApiKey:       os.Getenv("X_API_KEY"),
			ApiKeyPublic: os.Getenv("X_API_KEY_PUBLIC"),
			JWTSecret:    os.Getenv("JWT_SECRET"),
			NatsUrl:      os.Getenv("NATS_URL"),
		})

	s.Start(func(s *server.Server) {
		func(s *server.Server) {
			pe := event.NewPostEvent(s.Nats())
			pe.OnPublishCreated(func(post model.Post) {
				s.Hub(C.WS_CHANNEL_POST).Broadcast(
					util.NewWSBody(util.Body{
						Message: "post created",
						Data:    post,
					}, "[POST] CREATED"),
					nil)
			})
		}(s)

		func(s *server.Server) {
			app := s.App()
			app.Use(middleware.ApiKeyMiddleware(s.Config().ApiKeyPublic))
			app.Use(middleware.AuthMiddleware(s))
			wsRouter := app.Group("/ws")
			wsRouter.Use(middleware.VerifySocketUpgrade)
			wsRouter.Get("", websocket.New(func(cn *websocket.Conn) {
				user := cn.Locals("user").(*model.User)
				client := ws.NewClient(cn, *user)
				hub := s.Hub(C.WS_CHANNEL_POST)
				hub.OnConnect(client)
				defer hub.OnDisconnect(client)
				for {
					_, message, err := cn.ReadMessage()
					if err != nil {
						log.Println("read:", err)
						break
					}
					fmt.Printf("Message received: %s\n", message)
					hub.Broadcast(message, nil)
				}

			}))
		}(s)

	})
	s.Shutdown()
}
