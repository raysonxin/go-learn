package convert

import (
	"bytes"
	"encoding/binary"
	"math"
)

//字节转换成整形
func BytesToInt(b []byte) uint {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return uint(x)
}

func BytesToInt32(b []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return x
}

func BytesToUint16(b []byte) uint16 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint16
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return x
}

func ByteToBool(i byte) bool {
	if i == 1 {
		return true
	}
	return false
}

func Int16ToBytes(i uint) []byte {
	var buf = make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(i))
	return buf
}

func Int32ToBytes(i uint) []byte {
	var buf = make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(i))
	return buf
}

func Float64frombytes(bytes []byte) float64 {
    bits := binary.LittleEndian.Uint64(bytes)
    float := math.Float64frombits(bits)
    return float
}

func Float32bytes(float float32) []byte {
    bits := math.Float32bits(float)
    bytes := make([]byte, 4)
    binary.LittleEndian.PutUint32(bytes, bits)
    return bytes
}
