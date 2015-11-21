package rabbitmq

import(
	"fmt"
	"github.com/streadway/amqp"
)

type AmqpProducer struct{
	pubChan			chan AmqpMessage
}

//"amqp://guest:guest@localhost:5672/"

func (prod *AmqpProducer) StartProducer(uri string)(err error){
	connection,err:=amqp.Dial(uri)
	if err!=nil{
		return fmt.Errorf("Dial: %s",err)
	}
	defer connection.Close()
	
	channel,err:=connection.Channel()
	if err!=nil{
		return fmt.Errorf("Channel: %s",err)
	}
	prod.pubChan=make(chan AmqpMessage,10)
	
	for{
		select {
			case msg:=<- prod.pubChan:
			if err=channel.Publish(
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

func (prod *AmqpProducer) Publish(msg AmqpMessage){
	prod.pubChan<-msg
}