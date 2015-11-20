// MyTcpServer project main.go
package main

import (
	
	"atisafe/tcp"
)

func main() {
	
	/*test:="hello"+" world"
	fmt.Println(test)
	
	s:="df[jhj]dfd[l=10&s=23&c=12&f=110106000001&t=11010601]helloworldsdfdfdf[l=10&s=23&c=12&f=110106000001&t=11010601]helloworld"
	remain,pks:=tcp.CheckPacket(s)
	
	for k,v:=range pks{
		fmt.Println(k,v.Json())
	}
	
	fmt.Println("remain:",remain)*/
	

	
	port:="8080"
	tcp.StartListen(port)
	
}

