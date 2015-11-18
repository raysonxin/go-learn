package tcp

import (
	"swk/base"
)

type PackageData struct {
	Length      uint   //4字节 数据长度
	IsReply     bool   //1字节 是否回复
	CommandType int    //1字节 命令类型
	Command     int    //1字节 命令
	Identify    uint   //2字节 包标识
	Datas       []byte //n字节 数据正文
}

func NewPackageData(datas []byte) *PackageData {
	packageData := PackageData{}
	packageData.Length = convert.BytesToInt(datas[0:4])
	if int(datas[4]) == 1 {
		packageData.IsReply = true
	} else {
		packageData.IsReply = false
	}
	packageData.CommandType = int(datas[5])
	packageData.Command = int(datas[6])
	packageData.Identify = uint(convert.BytesToUint16(datas[7:9]))
	packageData.Datas = datas[9:]
	return &packageData
}
