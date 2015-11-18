package main

import (
	"encoding/json"
	"net"
	"swk/socket/tcp"
)

type Account struct {
	request *tcp.Request
	conn    net.Conn
}

func (a Account) Handle() {
	switch a.request.Command {
	case 0:
		a.Login(a.request.Datas)
		break
	}
}

func (a Account) Login(datas []byte) {
	account := AccountInfo{}
	json.Unmarshal(datas, &account)
	username := account.UserName
	clients[username] = a.conn
	result := a.request.ReplyPackage(a.request.Identify, nil, true)
	a.conn.Write(result)
}
