package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"

	"swk/inihelper"
)

var dbContext DbContext

func main() {
	fmt.Println("初始化")
	conf := inihelper.SetConfig("./conf.ini")
	url := conf.GetValue("Local", "Port")
	dbIp := conf.GetValue("DataBase", "Ip")
	fmt.Println("连接数据库")
	dbContext = NewDbContext(dbIp)

	fmt.Println("开启服务")
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transportFactory := thrift.NewTTransportFactory()
	var transport thrift.TServerTransport
	transport, _ = thrift.NewTServerSocket(":" + url)
	handler := DbServiceImpl{}
	processor := NewDbServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("服务开启成功,端口:", url)
	server.Serve()
}
