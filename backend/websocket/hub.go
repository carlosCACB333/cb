package ws

import (
	"fmt"
)

type Hub struct {
	clients map[string]*Client
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string]*Client),
	}
}

func (hub *Hub) OnConnect(client *Client) {
	fmt.Printf("Client %s connected\n ", client.ID)
	hub.clients[client.ID] = client

}

func (hub *Hub) OnDisconnect(client *Client) {
	fmt.Printf("Client %s disconnected", client.ID)
	client.Close()
	delete(hub.clients, client.ID)
}

func (hub *Hub) Broadcast(message any, exclude *Client) {
	for _, client := range hub.clients {
		if client != exclude {
			client.socket.WriteJSON(message)
		}
	}
}

func (hub *Hub) Close() {
	for _, client := range hub.clients {
		client.Close()
	}
}
