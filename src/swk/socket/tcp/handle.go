package tcp

import (
	"net"
	"time"
)

type Client struct {
	Identify interface{}
}

func (c *Client) Handle(conn *net.TCPConn, readerChannel chan []byte, closed func(conn *net.TCPConn)) {
	tmpBuffer := make([]byte, 0)
	buffer := make([]byte, 1024*1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			closed(conn)
			return
		}
		conn.SetReadDeadline(time.Now().Add(1 * time.Minute))
		tmpBuffer = UnPackage(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
}
