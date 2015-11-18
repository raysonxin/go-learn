package main

import (
	"net"
)

type Client struct {
	Conn *net.TCPConn
	Clients []string
    UserId string
}

