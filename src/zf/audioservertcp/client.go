package main

import (
	"net"
	"time"
)

type Client struct {
	Addr    *net.UDPAddr
	Clients []string
	Timer   *time.Timer
	Name    string
}

func NewClient(addr *net.UDPAddr, clients []string, name string) *Client {
	client := Client{}
	client.Addr = addr
	client.Clients = clients
	client.Name = name
	return &client
}
