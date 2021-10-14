package process

import (
	"fmt"
)

var userMgr *UserMgr

// UserMgr 用户管理类, 定义了在向用户表与相应的方法
type UserMgr struct {
	onlineUsers map[string]*UserProcess
}

func init() {
	userMgr = &UserMgr{onlineUsers: make(map[string]*UserProcess, 1024)}
}
func (receiver *UserMgr) AddOnlineUser(up *UserProcess) {
	receiver.onlineUsers[up.UserMobile] = up
}
func (receiver *UserMgr) DeleteOnlineUser(userMobile string) {
	delete(receiver.onlineUsers, userMobile)
}

func (receiver *UserMgr) GetOnlineUsers() map[string]*UserProcess {
	return receiver.onlineUsers
}
func (receiver *UserMgr) GetOnlineUserById(userMobile string) (up *UserProcess, err error) {
	up, ok := receiver.onlineUsers[userMobile]
	if !ok {
		// 不在线
		err = fmt.Errorf("用户%d 不存在", userMobile)
	}
	return
}
