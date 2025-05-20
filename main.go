package main

import (
	chat "chat-server/server"
)

func main() {
	chat.StartServer("localhost:9000")
}
