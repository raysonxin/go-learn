package udp

import (
	"bytes"
	"encoding/json"
	"swk/base"
)

type PackageHelper struct {
}

func (p *PackageHelper) Packet(isReply bool, identify uint32,
	number uint16, commandType uint8, command uint8, isSuccess bool, datas interface{}) []byte {
	data := bytes.Buffer{}
	if isReply {
		data.WriteByte(1)
	} else {
		data.WriteByte(0)
	}
	data.Write(convert.Int32ToBytes(uint(identify)))
	data.Write(convert.Int16ToBytes(uint(number)))
	data.WriteByte(byte(commandType))
	data.WriteByte(byte(command))
	if isSuccess {
		data.WriteByte(1)
	} else {
		data.WriteByte(0)
	}
	content := make([]byte, 0)
	if datas != nil {
		switch datas.(type) {
		case string:
			dataString, _ := datas.(string)
			content = []byte(dataString)
			break
		case []byte:
			content, _ = datas.([]byte)
		default:
			var err error
			content, err = json.Marshal(datas)
			if err != nil {
				return nil
			}
		}
	}
	data.Write(content) //写入正文
	return data.Bytes()
}

func (p *PackageHelper) UnPacket(datas []byte) *PackageData {
	packageData := &PackageData{}
	packageData.IsReply = convert.ByteToBool(datas[0])
	packageData.Identify = convert.BytesToInt32(datas[1:5])
	packageData.Number = convert.BytesToUint16(datas[5:7])
	packageData.CommandType = uint8(datas[7])
	packageData.Command = uint8(datas[8])
	packageData.IsSuccess = convert.ByteToBool(datas[9])
	packageData.Datas = datas[10:]
	return packageData
}
