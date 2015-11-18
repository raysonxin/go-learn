package tcp

import (
	"bytes"
	"encoding/json"
	"swk/base"
)

var PackageLength int = 13 //除数据正文以外包长度
var HeaderLength int = 4
var DataLength int = 4
var Header string = "swkz"

func UnPackage(buffer []byte, readerChannel chan []byte) []byte {
	length := len(buffer)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+PackageLength { //判断包长度
			break
		}
		headerPosition := i + HeaderLength              //头索引
		if string(buffer[i:headerPosition]) == Header { //找到头标识
			messageLength := convert.BytesToInt(buffer[headerPosition : headerPosition+DataLength]) //确定正文长度
			totalLength := i + PackageLength + int(messageLength)
			if length < totalLength { //判断长度
				break
			}
			data := buffer[headerPosition:totalLength] //取数据
			readerChannel <- data
			i = totalLength - 1
		}
	}
	if i == length {
		return make([]byte, 0)
	}
	return buffer[i:]
}

func ReplyPackage(datas interface{}, isReply bool, commandType int, command int, identify uint, isSuccess bool) []byte {
	data := bytes.Buffer{}
	data.Write([]byte(Header)) //写入头

	content := make([]byte, 0)
	if datas != nil {
		switch datas.(type) {
		case string:
			dataString, _ := datas.(string)
			content = []byte(dataString)
			break
		case []byte:
			content, _ = datas.([]byte)
		case int:
			dataInt, _ := datas.(int)
			content = convert.Int32ToBytes(uint(dataInt))
			break
		case float32:
			dataFloat, _ := datas.(float32)
			content = convert.Float32bytes(dataFloat)
		default:
			var err error
			content, err = json.Marshal(datas)
			if err != nil {
				return nil
			}
			break
		}
	}
	data.Write(convert.Int32ToBytes(uint(len(content)))) //正文长度
	if isReply {
		data.WriteByte(1)
	} else {
		data.WriteByte(0)
	}
	data.WriteByte(byte(commandType))
	data.WriteByte(byte(command))
	data.Write(convert.Int16ToBytes(identify))
	if isSuccess {
		data.WriteByte(1)
	} else {
		data.WriteByte(0)
	}

	data.Write(content) //写入正文
	return data.Bytes()
}
