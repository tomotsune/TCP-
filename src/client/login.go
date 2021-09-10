package client

import (
	"awesomeProject/src/common"
	"encoding/json"
	"errors"
	"net"
)

func Login(mobile, pwd string) (err error) {
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
	transfer := common.Transfer{Conn: conn}
	err = transfer.WritePkg(msg)
	if err != nil {
		return err
	}
	pkg, err := transfer.ReadPkg()
	if err != nil {
		return
	}
	res := common.LoginRes{}
	err = json.Unmarshal([]byte(pkg.Data), &res)
	if err != nil {
		return
	}
	if res.Code != 200 {
		err = errors.New(res.Error)
	}
	return
}
