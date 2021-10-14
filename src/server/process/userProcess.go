package process

import (
	common "awesomeProject/src/common"
	model2 "awesomeProject/src/common/model"
	"awesomeProject/src/server/model"
	"net"
)

// UserProcess 用户处理
type UserProcess struct {
	MemberDao  *model.MemberDao
	Conn       net.Conn
	UserMobile string
}

func NewUserProcess(conn net.Conn) *UserProcess {
	return &UserProcess{MemberDao: model.NewMemberDao(), Conn: conn}
}

func (receiver *UserProcess) Login(member *model2.Member) (err error) {
	// 注册
	err = receiver.MemberDao.Login(member)

	// 响应信息
	res := common.R{}
	if err != nil {
		res = common.R{Code: "500", Msg: err.Error()}
	} else {
		res = common.R{Code: "200"}
		// 加入用户列表
		receiver.UserMobile = member.Mobile
		userMgr.AddOnlineUser(receiver)
		// 广播上线通知
		err = receiver.NotifyOtherOnlineUser(
			&common.NotifyUserStatusMes{
				Status:     common.OnLine,
				UserMobile: member.Mobile})
	}
	// 3. 发送相应信息
	// 放入用户列表
	var users []string
	for _, v := range userMgr.onlineUsers {
		users = append(users, v.UserMobile)
	}
	res.Data = users
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
	// 发送响应信息
	msg := common.Message{Type: common.Res, Data: res}
	transfer := common.Transfer{Conn: receiver.Conn}
	err = transfer.WritePkg(&msg)
	return
}

func (receiver *UserProcess) NotifyOtherOnlineUser(mus *common.NotifyUserStatusMes) (err error) {
	// 通知某个用户
	notify := func(conn net.Conn) (err error) {
		msg := common.Message{
			Type: common.NotifyUserStatus,
			Data: *mus}
		transfer := common.Transfer{Conn: conn}
		err = transfer.WritePkg(&msg)
		return
	}
	// 遍历onlineUser, 逐个发送
	for _, v := range userMgr.onlineUsers {
		if v.UserMobile == mus.UserMobile {
			continue
		}
		err = notify(v.Conn)
		if err != nil {
			return
		}
	}
	return
}
