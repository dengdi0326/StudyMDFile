package main

import (
	"github.com/gorilla/websocket"
	"time"
)

type client struct {
	socket   *websocket.Conn
	send     chan *message
	room     *room
	userDate map[string]interface{}
}

func(c *client) read() {
	defer client.socket.Close()

	for {
		var msg *message
		err := client.socket.ReadJSON(&msg)
		if err != nil {
			return
		}

		msg.When = time.Now()
		msg.Name = c.userDate["name"].(string)

		msg.AvatarURL, _ = c.room.avatar.GetAvatarURL(c)

		c.room.forward<- msg
	}
}

func(c *client) write() {
	defer client.socket.Close()

	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
				return
		}
	}
}
