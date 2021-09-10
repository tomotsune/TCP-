package main

import (
	"awesomeProject/src/common"
	process2 "awesomeProject/src/server/process"
	"errors"
	"fmt"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (receiver *Processor) process() (err error) {
	transfer := common.Transfer{Conn: receiver.Conn}
	msg, err := transfer.ReadPkg()
	if err != nil {
		return
	}
	err = receiver.serverProcessMes(&msg)
	if err != nil {
		return
	}

	err = errors.New(receiver.Conn.RemoteAddr().String() + "断开连接")
	return
}
func (receiver *Processor) serverProcessMes(msg *common.Message) (err error) {
	userProcess := process2.UserProcess{Conn: receiver.Conn}
	switch msg.Type {
	case common.LoginReqType:
		err = userProcess.ServerProcessLogin(msg)
	//处理登录
	default:
		fmt.Println("消息类型不存在, 无法处理")
	}
	return
}
