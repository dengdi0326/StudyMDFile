package main

import (
	"log"
	"github.com/sapxry/ex3/trace"
	"github.com/gorilla/websocket"
	"net/http"
)

type room struct {
	forward chan []byte

	join chan *client

	leave chan *client

	clients map[*client]bool

	tracer trace.Tracer
}





func (r *room) Run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("New client joined")

		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left")

		case msg := <-r.forward:
			r.tracer.Trace("Message received", string(msg))

			for client := range r.clients {
				client.send <- msg
				r.tracer.Trace(" send to client")
			}
		}
	}
}

const(
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{WriteBufferSize:socketBufferSize, ReadBufferSize:socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join<-client
	defer func() {r.leave<-client}()
	go client.write()
	client.read()
}