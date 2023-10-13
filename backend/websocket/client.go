package ws

import (
	"github.com/carlosCACB333/cb-back/model"

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	model.User
	socket *websocket.Conn
}

func NewClient(socket *websocket.Conn, user model.User) *Client {
	return &Client{
		User:   user,
		socket: socket,
	}
}

func (c *Client) Close() {
	c.socket.Close()
}
