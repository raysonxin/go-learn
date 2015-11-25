// MyTcpServer project main.go
package main

import (
	"time"
	"fmt"
	"encoding/json"
	"atisafe/rabbitmq"
)

func main() {
	mgr:=rabbitmq.MqManager{
		CfgFile:"config.xml",
	}
	
	config,err:=mgr.Initialize()
	
	if err != nil{
		panic(err.Error())
	}
	
	fmt.Println(config)
	
	go mgr.StartProducer()
	
	mgr.Subscribe(config.Consumes,handMessage)
	
	for i:=1;i<100;i++{
		data:=rabbitmq.MsgData{
			MsgType:		"type",
			MsgBody:		"helloworld",
			MsgTime:		time.Now(),
		}
		
		msg:=rabbitmq.AmqpMessage{
			Exchange:		"exc1",
			RoutingKey:		"abc.123",
			Data:			data,
		}
		mgr.Publish(msg)
		
		time.Sleep(2000)
	}
}

func handMessage(msg rabbitmq.MsgData){
	if b,err:=json.Marshal(msg); err == nil{
		fmt.Println(string(b))
	}
}
