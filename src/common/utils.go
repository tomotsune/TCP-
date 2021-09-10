package common

import (
	"encoding/binary"
	"encoding/json"
	"net"
)

func ReadPkg(conn net.Conn) (mes Message, err error) {
	buf := make([]byte, 8049)
	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}
	pkgLen := int(binary.BigEndian.Uint32(buf[:4]))
	n, err := conn.Read(buf)
	if n != pkgLen || err != nil {
		return
	}
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		return
	}
	return
}
func WritePkg(conn net.Conn, msg []byte) (err error) {
	// 先发送一个长度
	var pkgLen [4]byte
	binary.BigEndian.PutUint32(pkgLen[:], uint32(len(msg)))
	n, err := conn.Write(pkgLen[:])
	if n != 4 || err != nil {
		return
	}
	n, err = conn.Write(msg)
	if n != len(msg) || err != nil {
		return
	}
	return
}
