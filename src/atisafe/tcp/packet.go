package tcp

type MyPacket struct{
	Length			int
	Sequence		int
	Command			int
	From			string
	To				string
	Data			string
}