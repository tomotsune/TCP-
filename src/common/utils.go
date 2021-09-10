package common

import (
	"encoding/binary"
	"encoding/json"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (receiver *Transfer) ReadPkg() (mes Message, err error) {
	//buf := make([]byte, 8049)
	_, err = receiver.Conn.Read(receiver.Buf[:4])
	if err != nil {
		return
	}
	pkgLen := int(binary.BigEndian.Uint32(receiver.Buf[:4]))
	n, err := receiver.Conn.Read(receiver.Buf[:])
	if n != pkgLen || err != nil {
		return
	}
	err = json.Unmarshal(receiver.Buf[:pkgLen], &mes)
	if err != nil {
		return
	}
	return
}
func (receiver *Transfer) WritePkg(msg []byte) (err error) {
	// 先发送一个长度
	// var pkgLen [4]byte
	binary.BigEndian.PutUint32(receiver.Buf[:4], uint32(len(msg)))
	n, err := receiver.Conn.Write(receiver.Buf[:4])
	if n != 4 || err != nil {
		return
	}
	n, err = receiver.Conn.Write(msg)
	if n != len(msg) || err != nil {
		return
	}
	return
}
