package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"swk/inihelper"
	"swk/socket/tcp"
)

var onlineClients map[string]*Client

func main() {
	onlineClients = make(map[string]*Client, 0)
	runtime.GOMAXPROCS(runtime.NumCPU())
	tcp.Header = "zfty"
	conf := inihelper.SetConfig("./conf.ini")
	port := conf.GetValue("Local", "Port")
	service := ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("服务开启成功")
	checkError(err)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		fmt.Println("客户端连接：", conn.RemoteAddr().String())
		go handleClient(conn)
	}
}

func handleClient(conn *net.TCPConn) {
	conn.SetNoDelay(true)
	client := tcp.Client{}
	readerChannel := make(chan []byte, 16)
	go handle(readerChannel, conn)
	client.Handle(conn, readerChannel, clientClose)
}

func clientClose(conn *net.TCPConn) {
	fmt.Println(conn.LocalAddr().String(),"关闭")
}

func handle(readerChannel chan []byte, conn *net.TCPConn) {
	client := &Client{conn, make([]string, 0), ""}
	for {
		select {
		case data := <-readerChannel:
			packageData := tcp.NewPackageData(data)
			switch packageData.CommandType {
			case 0:
				audioService := AudioService{packageData, client}
				audioService.Handle()
				break
			}
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
