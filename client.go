package main

import (
	"github.com/gorilla/websocket"
)

// Client represent a single chatting user

type client struct {

	// Socket is the web socket for this client
	socket *websocket.Conn

	// Receive is a channel to receive messages from others clients
	receive chan []byte

	// Room is the room this client is chatting in
	room *room
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
