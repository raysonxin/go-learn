package tcp

import(
	"encoding/json"
)

type MyPacket struct{
	Length			int
	Sequence		int
	Command			uint16
	FuncCode		byte
	OptCode			byte
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