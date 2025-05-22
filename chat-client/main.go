package main

import "chat-client/client"

func main() {
	client.StartClient("chat-server:9000")
}
