package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type MqManager struct{
	CfgFile				string
	Conn				*amqp.Connection
	Channel				*amqp.Channel
}

//load config.ini information
func (mgr *MqManager) Initialize(){
	uri:="amqp://guest:guest@localhost:5672/"
	connection,err:=amqp.Dial(uri)
	if err!=nil{
		fmt.Errorf("Dial: %s",err)
		return
	}	
	mgr.Conn=connection;
	
	channel,err:=connection.Channel()
	if err!=nil{
		fmt.Errorf("Channel: %s",err)
		return
	}
	mgr.Channel=channel
	
	excs,ques,binds:=GetConfigs()
	mgr.SetupExchanges(excs)
	mgr.SetupQueues(ques)
	mgr.SetBinds(binds)
}

func GetConfigs()([]ExchangeInfo,[]QueueInfo,[]BindInfo){
	//TODO::to load config.ini 
	return nil,nil,nil
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

func (mgr *MqManager) SetupQueues(ques []QueueInfo){
	for _,cfg:=range ques{
		if _,err:=mgr.Channel.QueueDeclare(
			cfg.Name,
			cfg.Durable,
			cfg.AutoDelete,
			cfg.Exclusive,
			cfg.IsNoWait,
			nil,
		);err !=nil{
			fmt.Println("Setup queue error")
		}
	}
}

func (mgr *MqManager) SetBinds(binds []BindInfo){
	for _,cfg:=range binds{
		if err:=mgr.Channel.ExchangeBind(
			cfg.QueueName,
			cfg.BindKey,
			cfg.ExcName,
			cfg.IsNoWait,
			nil,
		);err != nil{
			fmt.Println("set binding error")
		}
	}
}