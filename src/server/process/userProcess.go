package process

import (
	"awesomeProject/src/common"
	"encoding/json"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (receiver *UserProcess) ServerProcessLogin(msg *common.Message) (err error) {
	// 1. 先从msg中取出data, 并直接反序列化成LoginMsg
	var loginReq common.LoginReq
	err = json.Unmarshal([]byte(msg.Data), &loginReq)
	if err != nil {
		return
	}
	// 2. 验证
	ResMsg := common.Message{Type: common.LoginResType}
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
	ResMsg.Data = string(loginRes)
	res, err := json.Marshal(ResMsg)
	if err != nil {
		return
	}
	transfer := common.Transfer{Conn: receiver.Conn}
	err = transfer.WritePkg(res)
	if err != nil {
		return
	}
	return
}
