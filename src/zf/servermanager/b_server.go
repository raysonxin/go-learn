package main

import (
	"fmt"
	"net"
	"swk/socket/tcp"
)

type Server struct {
	Request *tcp.Request
	Conn    net.Conn
}

func (s Server) Handle() {
	switch s.Request.Command {
	case 0:
		s.GetServerInfos()
		break
	case 1:
		s.AddServerInfo()
		break
	}
}

func (s Server) AddServerInfo() {
	serverInfo := string(s.Request.Datas)
	_, err := redisHelper.LPush("Servers", serverInfo)
	if err != nil {
		fmt.Println(err)
		result := s.Request.ReplyPackage(s.Request.Identify, nil, false)
		s.Conn.Write(result)
	} else {
		result := s.Request.ReplyPackage(s.Request.Identify, nil, true)
		s.Conn.Write(result)
	}
}

func (s Server) GetServerInfos() {
	length, _ := redisHelper.LLen("Servers")
	jsonServers, _ := redisHelper.LRange("Servers", 0, length)
	result := s.Request.ReplyPackage(s.Request.Identify, jsonServers, true)
	s.Conn.Write(result)
}
