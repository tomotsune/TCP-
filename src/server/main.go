package main

import (
	"awesomeProject/src/common"
	binary "encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("listen 8889...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go process(conn)
	}
}
func process(conn net.Conn) {
	defer conn.Close()

	for {
		msg, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println(conn.RemoteAddr(), "断开链接")
				return
			}
			fmt.Println("读取失败, err=", err)
		}
		fmt.Println("msg=", msg)
	}
}
func readPkg(conn net.Conn) (mes common.Message, err error) {
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
