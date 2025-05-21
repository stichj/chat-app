package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func StartServer(address string, broker *Broker) {
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
			fmt.Println("Connection failed:", err)
			continue
		}

		go handleConnection(conn, broker)
	}
}

func handleConnection(conn net.Conn, broker *Broker) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	writer.WriteString("Please enter your username: \n")
	writer.Flush()
	name, err := reader.ReadString('\n')
	if err != nil {
		conn.Close()
		fmt.Println("Error reading username", err)
		return
	}

	name = strings.TrimSpace(name)

	client := &Client{
		Conn:     conn,
		Name:     name,
		Outbound: make(chan string),
	}

	broker.Register <- client

	go func() {
		defer conn.Close()
		for msg := range client.Outbound {
			_, err := fmt.Fprintln(client.Conn, msg)
			if err != nil {
				return
			}
		}
	}()

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		select {
		case broker.Broadcast <- fmt.Sprintf("[%s]: %s", client.Name, strings.TrimSpace(msg)):
		default:
			fmt.Println("Blocked! Could not send to broker")
		}
	}

	broker.Unregister <- client
	conn.Close()
}
