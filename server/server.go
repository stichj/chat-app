package chat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func StartServer(address string) {

	fmt.Printf("Starting chat server on %s ...\n", address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Errorf("Error creating TCP socket on %s: %v\n", address, err)
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
	defer conn.Close()
	_, err := fmt.Fprintf(conn, "Please enter your username: ")
	if err != nil {
		fmt.Println("Error writing username prompt", err)
	}

	reader := bufio.NewReader(conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username", err)
	}

	name = strings.TrimSpace(name)
	fmt.Printf("%s has joined!\n", name)
	fmt.Fprintf(conn, "Welcome to the chat!\n")
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%s disconnected.\n", name)
			return
		}
		fmt.Printf("[%s]: %s", name, msg)
	}
}
