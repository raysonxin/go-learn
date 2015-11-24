// MyTcpServer project main.go
package main

import (
	"io/ioutil"
	"fmt"
	"atisafe/rabbitmq"
	//"atisafe/tcp"
	"encoding/xml"
)

func main() {
	
	/*test:="hello"+" world"
	fmt.Println(test)
	
	s:="df[jhj]dfd[l=10&s=23&c=12&f=110106000001&t=11010601]helloworldsdfdfdf[l=10&s=23&c=12&f=110106000001&t=11010601]helloworld"
	remain,pks:=tcp.CheckPacket(s)
	
	for k,v:=range pks{
		fmt.Println(k,v.Json())
	}
	
	fmt.Println("remain:",remain)
	
	cfg:=rabbitmq.MqConfig{
		FilePath:"./config.ini",
	}
	
	configs,_:=cfg.LoadConfig()
	
	for _,v := range configs{
		for key,value:= range v{
			for ck,cv:=range value{
				fmt.Println(key,ck,cv)
			}
		}
	}*/
	
	content,err:=ioutil.ReadFile("config.xml")
	if err != nil{
		panic(err.Error())
	}

	var config AppConfig
	
	
	//port:="8080"
	//tcp.StartListen(port)
	
}

