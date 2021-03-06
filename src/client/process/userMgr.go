package process

import (
	"awesomeProject/src/common"
	"fmt"
	"net"
)

var (
	onlineUsers   []string
	CurConn       net.Conn
	CurUserMobile string
)

/**
登录提醒, 加入在线用户表
*/
func updateUserStatus(msg *common.Message) {
	userMobile := common.Converter(msg, "userMobile").(string)
	status := int(common.Converter(msg, "status").(float64))
	if status == common.OnLine {
		onlineUsers = append(onlineUsers, userMobile)
		fmt.Printf("-->  用户%v已登录\n", userMobile)
	} else {
		for i, v := range onlineUsers {
			if v == userMobile {
				onlineUsers = append(onlineUsers[:i], onlineUsers[:i+1]...)
			}
		}
	}
}

/**
显示在线用户
*/
func outputOnlineUser() {
	fmt.Println("当前在线用户")
	for _, user := range onlineUsers {
		fmt.Println("用户:", user)
	}
}
