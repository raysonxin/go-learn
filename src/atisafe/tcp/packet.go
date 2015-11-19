package tcp

import(
	"encoding/json"
)

type MyPacket struct{
	Length			int
	Sequence		int
	Command			int
	From			string
	To				string
	Data			string
}

//get the json string of MyPacket
func (pkt MyPacket) Json() (str string){
	if b,err:=json.Marshal(pkt);err==nil{
		return string(b)
	}
	return ""
}