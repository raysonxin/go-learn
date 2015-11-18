package udp

type PackageData struct {
	IsReply     bool   //是否回复 1字节
	Identify    uint32 //包标识  4字节
	Number      uint16 //包索引  2字节
	CommandType uint8  //命令类型 1字节
	Command     uint8  //命令     1字节
	IsSuccess   bool   //是否成功 1字节
	Datas       []byte //数据    n字节
}
