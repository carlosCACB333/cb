package handler

import (
	"cb/model"
	"cb/server"
	ws "cb/websocket"

	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

const (
	WS_TEST = "WS_TEST"
)

func Handler(s *server.Server) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		user := c.Locals("user").(*model.User)
		client := ws.NewClient(c, *user)
		hub := s.Hub(WS_TEST)

		hub.OnConnect(client)
		defer hub.OnDisconnect(client)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			fmt.Printf("Message received: %s\n", message)
			hub.Broadcast(message, nil)
		}

	})
}
