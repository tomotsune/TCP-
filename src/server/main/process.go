package main

import (
	"awesomeProject/src/common"
	"awesomeProject/src/common/model"
	process2 "awesomeProject/src/server/process"
	"fmt"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (receiver *Processor) process() (err error) {
	transfer := common.Transfer{Conn: receiver.Conn}
	for {
		msg, err := transfer.ReadPkg()
		if err != nil {
			return err
		}
		err = receiver.serverProcessMes(&msg)
		if err != nil {
			return err
		}
	}
}
func (receiver *Processor) serverProcessMes(msg *common.Message) (err error) {
	userProcess := process2.NewUserProcess(receiver.Conn)
	smsProcess := process2.SmsProcess{}
	switch msg.Type {
	case common.Login:
		err = userProcess.Login(&model.Member{Mobile: common.Converter(msg, "mobile").(string), Pwd: common.Converter(msg, "pwd").(string)})
	case common.Register:
		err = userProcess.Register(&model.Member{Mobile: common.Converter(msg, "mobile").(string), Pwd: common.Converter(msg, "pwd").(string)})
	case common.SMS:
		err = smsProcess.SendGroup(&common.SmsMsg{Content: common.Converter(msg, "content").(string), UserMobile: common.Converter(msg, "userMobile").(string)})
	default:
		fmt.Println("消息类型不存在, 无法处理")
	}
	return
}
