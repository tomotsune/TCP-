package process

import (
	common "awesomeProject/src/common"
	model2 "awesomeProject/src/common/model"
	"awesomeProject/src/server/model"
	"net"
)

type UserProcess struct {
	MemberDao *model.MemberDao
	Conn      net.Conn
}

func NewUserProcess(conn net.Conn) *UserProcess {
	return &UserProcess{MemberDao: model.NewMemberDao(), Conn: conn}
}

func (receiver *UserProcess) Login(member *model2.Member) (err error) {
	// 准备相应信息
	err = receiver.MemberDao.Login(member)
	res := common.R{}
	if err != nil {
		res = common.R{Code: "500", Msg: err.Error()}
	} else {
		res = common.R{Code: "200"}
	}
	// 3. 发送相应信息
	msg := common.Message{Type: common.Res, Data: res}
	transfer := common.Transfer{Conn: receiver.Conn}
	err = transfer.WritePkg(&msg)
	return
}
func (receiver *UserProcess) Register(member *model2.Member) (err error) {
	err = receiver.MemberDao.Register(member)
	res := common.R{}
	if err != nil {
		res = common.R{Code: "500", Msg: err.Error()}
	} else {
		res = common.R{Code: "200"}
	}
	if err != nil {
		return
	}
	// 3. 发送相应信息
	msg := common.Message{Type: common.Res, Data: res}
	transfer := common.Transfer{Conn: receiver.Conn}
	err = transfer.WritePkg(&msg)
	return
}
