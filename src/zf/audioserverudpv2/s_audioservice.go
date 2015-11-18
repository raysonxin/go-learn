package main

import (
	"fmt"
	"strings"
	"swk/socket/tcp"
)

type AudioService struct {
	Package *tcp.PackageData
	Client  *Client
}

func (a *AudioService) Handle() {
	switch a.Package.Command {
	case 0:
		a.login()
		break
	case 1:
		a.SendAudio()
		break
	case 2:
		a.RequestTalk()
		break
	}
}

func (a *AudioService) login() {
	a.Client.UserId = string(a.Package.Datas)
	onlineClients[a.Client.UserId] = a.Client
	fmt.Println(a.Client.UserId, "登陆")
	var result = tcp.ReplyPackage(nil, true, a.Package.CommandType, a.Package.Command, a.Package.Identify, true)
	a.Client.Conn.Write(result)
}

func (a *AudioService) RequestTalk() {
	users := strings.Split(string(a.Package.Datas), "|")
	a.Client.Clients = append(a.Client.Clients, users[0:]...)

	var result = tcp.ReplyPackage(nil, true, a.Package.CommandType, a.Package.Command, a.Package.Identify, true)
	a.Client.Conn.Write(result)
}

func (a *AudioService) SendAudio() {
	var result []byte
    result = tcp.ReplyPackage(a.Package.Datas, false, 0, 0, a.Package.Identify, true)
	for i := 0; i < len(a.Client.Clients); i++ {
		client, ok := onlineClients[a.Client.Clients[i]]
		if ok {
			client.Conn.Write(result)
		}
	}
}
