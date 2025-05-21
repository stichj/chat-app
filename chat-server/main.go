package main

import (
	"chat-server/server"
	"fmt"
)

func main() {
	fmt.Println("[Main] Starting server...")
	broker := server.NewBroker()
	go broker.Start()
	server.StartServer("localhost:9000", broker)
}
