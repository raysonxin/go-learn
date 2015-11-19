package tcp

import(
	"fmt"
	"net"
	"os"
)

func StartListen(port string){
	service:=":"+port
	tcpAddr,err:=net.ResolveTCPAddr("tcp4",service)
	listener,err:=net.ListenTCP("tcp",tcpAddr)
	if err!=nil{
		fmt.Printf("start tcp server error %s",err.Error())
		return;
	}
	
	fmt.Println("service is running...")
	for{
		conn,err:=listener.AcceptTCP()
		if err!=nil{
			continue
		}
		fmt.Println("accept tcp connection",conn.RemoteAddr().String())
		go handleClient(conn)
	}
}

func handleClient(conn *net.TCPConn){
	melonClient:=MelonClient{
		Conn:conn,
		Id:conn.RemoteAddr().String(),
	}
	readChan:=make(chan []MyPacket,16)
	go handlePacket(readChan,conn)
	melonClient.Run(conn,readChan,clientClosed)
}

func handlePacket(readChan chan []MyPacket,conn *net.TCPConn){
	for{
		select{
			case pkts:=<-readChan:
			for _,v :=range pkts{
				fmt.Println(v.Json())
			}
			break
		}
	}
}

func clientClosed(conn *net.TCPConn){
	
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}