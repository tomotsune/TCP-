package process

import (
	"awesomeProject/src/common"
	"awesomeProject/src/common/model"
	"errors"
	"net"
)

type UserProcess struct {
}

func (receiver *UserProcess) Login(member *model.Member) (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", "localhost:8889")
	if err != nil {
		return
	}
	transfer := common.Transfer{Conn: conn}
	err = transfer.WritePkg(&common.Message{Type: common.Login, Data: *member})
	if err != nil {
		return
	}
	res, err := transfer.ReadPkg()
	if err != nil {
		return
	}

	if common.Converter(&res, "code").(string) != "200" {
		err = errors.New(common.Converter(&res, "msg").(string))
		return
	}

	// 载入在线用户列表
	users := common.Converter(&res, "data").([]interface{})
	for _, user := range users {
		onlineUsers = append(onlineUsers, user.(string))
	}
	// 存储用户, 返回存储连接
	CurUserMobile = member.Mobile
	return
}

func (receiver *UserProcess) Register(member *model.Member) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	defer conn.Close()
	if err != nil {
		return
	}
	transfer := common.Transfer{Conn: conn}
	err = transfer.WritePkg(&common.Message{Type: common.Register, Data: *member})
	if err != nil {
		return
	}
	res, err := transfer.ReadPkg()
	if err != nil {
		return
	}
	if res.Data.(map[string]interface{})["code"].(string) != "200" {
		err = errors.New(res.Data.(map[string]interface{})["msg"].(string))
	}
	return
}
