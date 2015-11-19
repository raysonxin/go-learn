package tcp

import(
	"net"
	"time"
)

type MelonClient struct{
	Id			string
	Conn		*net.TCPConn
}

func (cli *MelonClient) Disconnect(){
	cli.Conn.Close()
}

func (cli *MelonClient) Run(conn *net.TCPConn,recvChan chan []byte,closed func(conn *net.TCPConn)){
	//tempBuffer:=make([]byte,0)
	buffer:=make([]byte,1024*1024)
	for{
		n,err:=conn.Read(buffer)
		if err!=nil{
			closed(conn)
			return
		}
		if n==0{
			closed(conn)
			return
		}
		conn.SetReadDeadline(time.Now().Add(1*time.Minute))
		//tempBuffer
	}
}