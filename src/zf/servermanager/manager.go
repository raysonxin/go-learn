package main

import (
	"fmt"
	"menteslibres.net/gosexy/redis"
	"net"
	"os"
	"strconv"
	"swk/inihelper"
	"swk/socket/tcp"
)

var clients map[string]net.Conn
var redisHelper *redis.Client

func main() {

	fmt.Println("连接Reids服务器")
	redisHelper = redis.New()
	conf := inihelper.SetConfig("./conf.ini")
	redisIp := conf.GetValue("Redis", "Ip")
	redisPort, _ := strconv.ParseInt(conf.GetValue("Redis", "Port"), 10, 32)
	err := redisHelper.Connect(redisIp, uint(redisPort))

	if err != nil {
		fmt.Println("连接失败: %s\n", err.Error())
		return
	}
	fmt.Println("Redis连接成功")

	fmt.Println("正在开启服务")
	clients = make(map[string]net.Conn, 0)
	service := ":8000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	fmt.Println("服务开启成功")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	client := tcp.Client{}
	readerChannel := make(chan []byte, 16)
	go handle(readerChannel, conn)
	client.Handle(conn, readerChannel)
}

func handle(readerChannel chan []byte, conn net.Conn) {
	for {
		select {
		case data := <-readerChannel:
			request := tcp.NewRequest(data)
			switch request.CommandType {
			case 0:
				account := Account{request, conn}
				account.Handle()
				break
			case 1:
				server := Server{request, conn}
				server.Handle()
				break
			}
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %s", err.Error())
	}
}
