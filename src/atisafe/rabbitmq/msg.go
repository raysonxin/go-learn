package rabbitmq

import(
	"time"
	"encoding/json"
)

type AmqpMessage struct{
	Exchange			string
	RoutingKey			string
	Data				MsgData
}

type MsgData struct{
	MsgType				string
	MsgBody				string
	MsgTime				time.Time
}

func (data MsgData) Json() (str string){
	if b,err:=json.Marshal(data);err==nil{
		return string(b)
	}
	return ""
}