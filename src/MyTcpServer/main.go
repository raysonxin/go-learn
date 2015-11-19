// MyTcpServer project main.go
package main

import (
	"fmt"
	"atisafe/tcp"
	"encoding/json"
)

func main() {
	s:="df[jhj]dfd[l=10&s=23&c=12&f=110106000001&t=11010601]helloworldsdfdfdf[l=10&s=23&c=12&f=110106000001&t=11010601]helloworld"
	remain,pks:=tcp.CheckPacket(s)
	
	for k,v:=range pks{
		if b,err:=json.Marshal(v);err==nil{
			fmt.Println(k,(string(b)))
		}
	}
	
	fmt.Println("remain:",remain)
}

