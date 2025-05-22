package client

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func StartClient(address string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Connect to TCP chat server
	conn, err := net.Dial("tcp", address)
	fmt.Println("Connecting to", address, "...")
	if err != nil {
		fmt.Printf("Error connecting to server %s via tcp: %s\n", address, err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connected to chat server")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		shutdown(conn, cancel)
	}()

	//Read welcome prompt (asking for username) from chat server
	welcomePrompt, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading welcome prompt from server:", err)
	}
	welcomePrompt = strings.TrimSpace(welcomePrompt)
	fmt.Println(welcomePrompt)

	//Read username input from stdin
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading username from stdin:", err)
	}
	fmt.Fprint(conn, username)
	fmt.Printf("Welcome %v! Start chatting below: \n", strings.TrimSpace(username))

	//Read user input and write message to server
	go sendMessages(ctx, conn, cancel)

	//Listen for server broadcasts
	broadcastReader := bufio.NewReader(conn)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nDisconnected from server.")
			conn.Close()
			return
		default:
			broadcast, err := broadcastReader.ReadString('\n')
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				fmt.Println("Error reading broadcast:", err)
				cancel()
				return
			}
			fmt.Println(strings.TrimSpace(broadcast))
		}
	}
}

func sendMessages(ctx context.Context, conn net.Conn, cancel context.CancelFunc) {
	reader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading your input:", err)
				continue
			}

			trimmed := strings.TrimSpace(msg)
			if trimmed == "/quit" || trimmed == "/q" {
				shutdown(conn, cancel)
				return
			}

			_, err = fmt.Fprintln(conn, msg)
			if err != nil {
				fmt.Println("Error sending your message:", err)
				continue
			}

			// Clean up echoed input
			fmt.Print("\033[1A") // move cursor up
			fmt.Print("\033[2K") // clear line
		}
	}
}

func shutdown(conn net.Conn, cancel context.CancelFunc) {
	fmt.Println("\nShutting down ...")
	fmt.Println("Leaving chat ...")
	fmt.Println("Goodbye!")
	cancel()
	conn.Close()
	os.Exit(0)
}
