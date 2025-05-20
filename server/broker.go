package server

import "fmt"

type Broker struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan string
}

func NewBroker() *Broker {
	return &Broker{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan string),
	}
}

func (b *Broker) Start() {
	for {
		select {
		case client := <-b.Register:
			b.Clients[client] = true
			go func(msg string) {
				b.Broadcast <- msg
			}(client.Name + " has joined the chat!")

		case client := <-b.Unregister:
			if _, ok := b.Clients[client]; ok {
				delete(b.Clients, client)
				close(client.Outbound)
				go func(msg string) {
					b.Broadcast <- msg
				}(client.Name + " has left the chat!")

			}
		case msg := <-b.Broadcast:
			fmt.Printf("Broadcasting: %s\n", msg)
			for client := range b.Clients {
				select {
				case client.Outbound <- msg:
				default:
					close(client.Outbound)
					delete(b.Clients, client)
				}
			}
		}
	}
}
