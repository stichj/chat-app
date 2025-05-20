package chat

import (
	"fmt"
	"net"
)

func StartServer(address string) {

	fmt.Printf("Starting chat server on %s ...\n", address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Errorf("Error creating TCP socket on %s!\n", address)
	}
	defer listener.Close()
	fmt.Printf("Server running on %s\n", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection failed", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Fprintf(conn, "Welcome to the chat!\n")
	conn.Close()
}
