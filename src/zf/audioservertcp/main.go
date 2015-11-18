package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"swk/inihelper"
	"swk/socket/udp"
	"swk/base"
	
)

var conn *net.UDPConn
var onlineClients map[string]*Client

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	conf := inihelper.SetConfig("./conf.ini")
	port := conf.GetValue("Local", "Port")
	service := ":" + port
	onlineClients = make(map[string]*Client, 0)
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	_conn, err := net.ListenUDP("udp", udpAddr)
	conn = _conn
	checkError(err)
	fmt.Println("服务开启成功：", service)
	for {
		handleClient()
	}
}

func handleClient() {
	var buf [512]byte
	count, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	go handle(buf[:count], addr)

}

var packageHelper udp.PackageHelper = udp.PackageHelper{}

func handle(datas []byte, addr *net.UDPAddr) {
	packageData := packageHelper.UnPacket(datas)
	switch packageData.Command {
	case 0:
		go login(packageData, addr)
		break
	case 1:
		go SendAudio(packageData, addr, GetClient(addr))
		break
	case 2:
		go RequestTalk(packageData, addr, GetClient(addr))
		break
	}
}

func GetClient(addr *net.UDPAddr) *Client {
	var client *Client
	for _, value := range onlineClients {
		if value.Addr.String() == addr.String() {
			client = value
			break
		}
	}
	return client
}


func login(packageData *udp.PackageData, addr *net.UDPAddr) {
	userName := string(packageData.Datas)
	fmt.Println(userName)
	onlineClients[userName] = NewClient(addr, make([]string, 0), userName)
	var result = packageHelper.Packet(true, packageData.Identify, packageData.Number, packageData.CommandType, packageData.Command, true, nil)
	conn.WriteToUDP(result, addr)
}

func RequestTalk(packageData *udp.PackageData, addr *net.UDPAddr, client *Client) {
	users := strings.Split(string(packageData.Datas), "|")
	for i := 0; i < len(users); i++ {
		client.Clients = append(client.Clients, users[i])
	}
	var result = packageHelper.Packet(true, packageData.Identify, packageData.Number, packageData.CommandType, packageData.Command, true, nil)
	conn.WriteToUDP(result, addr)
}

func SendAudio(packageData *udp.PackageData, addr *net.UDPAddr, client *Client) {	
	var result []byte
	if client == nil {
		return
	}
	if len(client.Clients) == 0 {
		result = packageHelper.Packet(false, packageData.Identify, 0, 0, 1, true, nil)
		conn.WriteToUDP(result, addr)
		return
	} else {
		result = packageHelper.Packet(false, packageData.Identify, 0, 0, 0, true, packageData.Datas)
	}
	fmt.Println("收到数据",client.Name,convert.BytesToInt32(packageData.Datas[0:4]))
	for i := 0; i < len(client.Clients); i++ {
		sendClient, ok := onlineClients[client.Clients[i]]
		if ok {
			conn.WriteToUDP(result, sendClient.Addr)
		}else{
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
	}
}
