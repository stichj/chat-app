package server

import "net"

type Client struct {
	Conn     net.Conn
	Name     string
	Outbound chan string
}
