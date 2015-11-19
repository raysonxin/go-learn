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
	melonClient:=MelonClient{}
	melonClient.Conn=conn
	melonClient.Id=conn.RemoteAddr().String()
	readChan:=make(chan []byte,16)

	melonClient.Run(conn,readChan,clientClosed)
}

func clientClosed(conn *net.TCPConn){
	
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}