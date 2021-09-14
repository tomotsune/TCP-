package process

import (
	"awesomeProject/src/common"
	"net"
)

type SmsProcess struct {
}

func (receiver *SmsProcess) SendGroup(sms *common.SmsMsg) (err error) {
	// 通知某个用户
	notify := func(conn net.Conn) (err error) {
		msg := common.Message{Type: common.SMS, Data: sms}
		transfer := common.Transfer{Conn: conn}
		err = transfer.WritePkg(&msg)
		return
	}
	// 遍历onlineUser, 逐个发送
	for _, v := range userMgr.onlineUsers {
		if v.UserMobile == sms.UserMobile {
			continue
		}
		err = notify(v.Conn)
		if err != nil {
			return
		}
	}
	return
}
