package main

import (
	"awesomeProject/src/common"
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
		msg, err := common.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println(conn.RemoteAddr(), "断开链接")
				return
			}
			fmt.Println("读取失败, err=", err)
			return
		}
		err = serverProcessMes(conn, &msg)
		if err != nil {
			fmt.Println("消息处理失败, err=", err)
			return
		}
	}
}

func serverProcessMes(conn net.Conn, msg *common.Message) (err error) {
	switch msg.Type {
	case common.LoginReqType:
		err = serverProcessLogin(conn, msg)
	//处理登录
	default:
		fmt.Println("消息类型不存在, 无法处理")
	}
	return
}
func serverProcessLogin(conn net.Conn, msg *common.Message) (err error) {
	// 1. 先从msg中取出data, 并直接反序列化成LoginMsg
	var loginReq common.LoginReq
	err = json.Unmarshal([]byte(msg.Data), &loginReq)
	if err != nil {
		return
	}
	// 2. 验证
	msg = &common.Message{Type: common.LoginResType}
	var loginRes []byte
	if loginReq.Mobile == "root" && loginReq.Pwd == "0000" {
		loginRes, err = json.Marshal(common.LoginRes{Code: 200})
		if err != nil {
			return
		}
	} else {
		loginRes, err = json.Marshal(common.LoginRes{Code: 500, Error: "登录信息错误"})
		if err != nil {
			return
		}
	}
	msg.Data = string(loginRes)
	res, err := json.Marshal(*msg)
	if err != nil {
		return
	}
	err = common.WritePkg(conn, res)
	if err != nil {
		return
	}
	return
}
