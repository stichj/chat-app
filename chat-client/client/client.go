package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func StartClient(address string) {
	//Connect to TCP chat server
	conn, err := net.Dial("tcp", address)
	fmt.Println("Connecting to", address, "...")
	if err != nil {
		fmt.Printf("Error connecting to server %s via tcp: %s\n", address, err)
	}
	defer conn.Close()
	fmt.Println("Connected to chat server")

	//Read welcome prompt (asking for username) from chat server
	welcomePrompt, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading welcome prompt from server:", err)
	}
	fmt.Println(welcomePrompt)

	//Read username input from stdin
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username from stdin:", err)
	}
	fmt.Fprint(conn, username)
}
