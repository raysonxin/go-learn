// MyTcpServer project main.go
package main

import (
	"fmt"
	"net"
)

func main() {
	listen,err:=net.Listen("tcp",":8080")
	if err!=nil{
		panic(err.Error())
	}
	
	for{
		conn,err:=listen.Accept()
		if err!=nil{
			panic(err.Error())
		}
		
		go handleConnection(conn)
	}
	fmt.Println("Hello World!")
}

func handleConnection(conn net.Conn){
	fmt.Println("accept a connection",conn.RemoteAddr())
	
}
