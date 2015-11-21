package rabbitmq

import (
	"fmt"
)

import(
	"github.com/streadway/amqp"
)

type MqManager struct{
	CfgFile				string
	Conn				*amqp.Connection
	Channel				*amqp.Channel
}

//load config.ini information
func (mgr *MqManager) Initialize(){
	
}

func (mgr *MqManager) SetupExchanges(excs []ExchangeInfo){
	for _,cfg:=range excs{
		if err:=mgr.Channel.ExchangeDeclare(
			cfg.Name,
			cfg.ExcType,
			cfg.Durable,
			cfg.AutoDelete,
			cfg.Internal,
			false,
			nil,
		);err!=nil{
			fmt.Println("setup exchange error")
		}
	}
}
