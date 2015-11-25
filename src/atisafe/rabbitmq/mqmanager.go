package rabbitmq

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"fmt"
	"github.com/streadway/amqp"
)

type MqManager struct{
	CfgFile				string
	Conn				*amqp.Connection
	Channel				*amqp.Channel
	pubChan				chan AmqpMessage
}

//load config.ini information
func (mgr *MqManager) Initialize()(c MqConfig,e error){
	
	config,err:=mgr.getConfigs()
	if err != nil{
		panic(err.Error())
	}
	
	connection,err:=amqp.Dial(config.ServerUri)
	if err!=nil{
		fmt.Errorf("Dial: %s",err)
		panic(err.Error())
	}	
	mgr.Conn=connection;
	
	channel,err:=connection.Channel()
	if err!=nil{
		fmt.Errorf("Channel: %s",err)
		panic(err.Error())
	}
	mgr.Channel=channel
	
	mgr.setupExchanges(config.Exchanges)
	mgr.setupQueues(config.Queues)
	mgr.setupBinds(config.Binds)
	
	return config,nil
}

//get config.xml information
func (mgr *MqManager) getConfigs()(MqConfig,error){
	
	content,err:=ioutil.ReadFile(mgr.CfgFile)
	if err!=nil{
		panic(err.Error())
	}
	
	var config MqConfig
	err=xml.Unmarshal(content,&config)
	return config,err
}

//setup exchanges at rabbitmq server
func (mgr *MqManager) setupExchanges(excs []Exchange){
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

//setup queue at rabbitmq server
func (mgr *MqManager) setupQueues(ques []Queue){
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

//setup bind at rabbitmq server
func (mgr *MqManager) setupBinds(binds []Bind){
	for _,cfg:=range binds{
		if err:=mgr.Channel.QueueBind(
			cfg.QueueName,
			cfg.BindKey,
			cfg.ExcName,
			cfg.IsNoWait,
			nil,
		);err != nil{
			fmt.Println("set binding error",err.Error())
		}
	}
}

//start producer task
func (mgr *MqManager) StartProducer()(err error){
	
	mgr.pubChan=make(chan AmqpMessage,10)
	
	for{
		select {
			case msg:=<- mgr.pubChan:
			if err=mgr.Channel.Publish(
				msg.Exchange,
				msg.RoutingKey,
				false,
				false,
				amqp.Publishing{
					Headers:			amqp.Table{},
					ContentEncoding:	"text/plain",
					Body:				[]byte(msg.Data.Json()),
					DeliveryMode:		amqp.Transient,
					Priority:			0,
				},
			);err !=nil{
				fmt.Println("publish error")
			}
			break
		}
	}
	
	return nil
}

//publish message
func (mgr *MqManager) Publish(msg AmqpMessage){
	mgr.pubChan<-msg
}

func (mgr *MqManager) Subscribe(csus []Consume,OnRecvMessage func(msg MsgData)){
	for _,c:= range csus{
		
		deliveries,err:=mgr.Channel.Consume(
			c.ListenQueue,
			c.ConsumerTag,
			false,
			false,
			false,
			false,
			nil,
		)
		
		if err != nil{
			fmt.Println("sub error",err.Error())
		}
		
		go handleRecv(deliveries,OnRecvMessage)
	}
}

func handleRecv(deliveries <-chan amqp.Delivery,OnRecvMessage func(msg MsgData)){
	for d:=range deliveries{
		var msg MsgData
		if err:=json.Unmarshal(d.Body,&msg);err ==nil{
			go OnRecvMessage(msg)
		}
	}
}