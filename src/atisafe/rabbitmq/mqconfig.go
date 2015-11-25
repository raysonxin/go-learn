package rabbitmq

import (
	"encoding/xml"
)

type MqConfig struct{
	XMLName			xml.Name `xml:"RabbitMQ"`
	ServerUri		string `xml:"uri"`
	Exchanges		[]Exchange `xml:"exchange"`
	Queues			[]Queue `xml:"queue"`
	Binds			[]Bind `xml:"bind"`
	Consumes		[]Consume `xml:"consume"`
}

type Exchange struct{
	Name			string `xml:"name,attr"`
	ExcType			string `xml:"type,attr"`
	Durable			bool `xml:"durable,attr"`
	AutoDelete		bool `xml:"autodelete,attr"`
	Internal		bool `xml:"internal,attr"`
	IsNoWait		bool `xml:"wait,attr"`
}

type Queue struct{
	Name			string `xml:"name,attr"`
	Durable			bool `xml:"durable,attr"`
	AutoDelete		bool `xml:"autodelete,attr"`
	Exclusive		bool `xml:"exclusive,attr"`
	IsNoWait		bool `xml:"wait,attr"`
}

type Bind struct{
	QueueName		string `xml:"queue,attr"`
	ExcName			string `xml:"exchange,attr"`
	BindKey			string `xml:"key,attr"`
	IsNoWait		bool `xml:"wait,attr"`
}

type Consume struct{
	ListenQueue		string `xml:"queue,attr"`
	ConsumerTag		string `xml:"tag,attr"`
}