package rabbitmq

import(
	"bufio"
	"io"
	"os"
	"strings"
)

type MqConfig struct{
	FilePath		string
	Configs			[]map[string]map[string]string
}

type AppConfig struct{
	Exchanges		[]ExchangeInfo
	Queues			[]QueueInfo
	Binds			[]BindInfo
}

type ExchangeInfo struct{
	Name			string `xml:"name,attr"`
	ExcType			string `xml:"type,attr"`
	Durable			bool `xml:"durable,attr"`
	AutoDelete		bool `xml:"autodelete,attr"`
	Internal		bool `xml:"internal,attr"`
	IsNoWait		bool `xml:"wait,attr"`
}

type QueueInfo struct{
	Name			string `xml:"name,attr"`
	Durable			bool `xml:"durable,attr"`
	AutoDelete		bool `xml:"autodelete,attr"`
	Exclusive		bool `xml:"exclusive,attr"`
	IsNoWait		bool `xml:"wait,attr"`
}

type BindInfo struct{
	QueueName		string `xml:"queue,attr"`
	ExcName			string `xml:"exchange,attr"`
	BindKey			string `xml:"key,attr"`
	IsNoWait		bool `xml:"wait,attr"`
}


func (c *MqConfig) Config(filepath string) {
	c.FilePath=filepath
}


func (c *MqConfig) LoadConfig() ([]map[string]map[string]string,error){
	file,err:=os.Open(c.FilePath)
	if err!=nil{
		return nil,err
	}
	defer file.Close()
	
	var data map[string]map[string]string
	var section string
	buf:=bufio.NewReader(file)
	
	for{
		l,err:=buf.ReadString('\n')
		line:=strings.TrimSpace(l)
		if err != nil{
			if err != io.EOF{
				return nil,err
			}
			if 0 ==len(line){
				break
			}
		}
		switch{
			case len(line)==0:
			case line[0]=='[' && line[len(line)-1]==']':
				section=strings.TrimSpace(line[1:len(line)-1])
				data=make(map[string]map[string]string)
				data[section]=make(map[string]string)
			default:
				i:=strings.IndexAny(line,"=")
				value:=strings.TrimSpace(line[i+1:len(line)])
				data[section][strings.TrimSpace(line[0:i])]=value
				if c.uniquappend(section)==true{
					c.Configs=append(c.Configs,data)
				}
		}
	}
	return c.Configs,nil
}

func (c *MqConfig) uniquappend(conf string) bool {
	for _, v := range c.Configs {
		for k, _ := range v {
			if k == conf {
				return false
			}
		}
	}
	return true
}