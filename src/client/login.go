package main

import (
	"awesomeProject/src/common"
	"encoding/binary"
	"encoding/json"
	"net"
)

func login(mobile, pwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	defer conn.Close()
	if err != nil {
		return
	}
	req, err := json.Marshal(common.LoginReq{Mobile: mobile, Pwd: pwd})
	if err != nil {
		return
	}
	msg, err := json.Marshal(common.Message{Type: common.LoginReqType, Data: string(req)})
	if err != nil {
		return
	}
	var pkgLen [4]byte
	binary.BigEndian.PutUint32(pkgLen[:], uint32(len(msg)))
	n, err := conn.Write(pkgLen[:])
	if n != 4 || err != nil {
		return
	}
	_, err = conn.Write(msg)
	if err != nil {
		return
	}
	// TODO: 处理服务器返回的消息

	return
}
