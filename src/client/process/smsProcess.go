package process

import (
	"awesomeProject/src/common"
	"fmt"
	"net"
)

type MsmProcess struct {
	conn net.Conn
}

func (receiver *MsmProcess) SendMsg(content string) (err error) {
	msm := common.SmsMsg{Content: content, UserMobile: CurUserMobile}
	msg := common.Message{Type: common.SMS, Data: msm}
	transfer := common.Transfer{Conn: receiver.conn}
	err = transfer.WritePkg(&msg)
	return
}
func (receiver *MsmProcess) ReceiveMsg(msg *common.Message) {
	userMobile := common.Converter(msg, "userMobile")
	content := common.Converter(msg, "content")
	fmt.Printf("用户%v->%v", userMobile, content)
	return
}
