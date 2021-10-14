// @Title  服务总控
// @Description
// @Author  haipinHu  08/10/2021 08:23
// @Update  haipinHu  08/10/2021 08:23
package process

import (
	"awesomeProject/src/common"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// 显示登录成功后的界面
func ShowMenu() {
	key := 0
	loop := true
	for loop {
		fmt.Println("----------登录成功---------")
		fmt.Println("\t\t\t1. 显示在线用户列表")
		fmt.Println("\t\t\t2. 发送消息")
		fmt.Println("\t\t\t3. 消息列表")
		fmt.Println("\t\t\t4. 退出系统")
		fmt.Println("请选择(1-4)")
		msmProcess := MsmProcess{conn: CurConn}
		fmt.Scanln(&key)
		switch key {
		case 1:
			outputOnlineUser()
		case 2:
			content := ""
			fmt.Print("_>")
			fmt.Scanln(&content)
			err := msmProcess.SendMsg(content)
			if err != nil {
				fmt.Println("err=", err)
			}
		case 4:
			// 不会执行defer
			os.Exit(0)
		default:
			fmt.Println("请选择(1-4)")
		}
	}
}

// 登录成功后. 该协程被启动, 监听所有消息
func ServerProcessMsg(conn net.Conn) {
	CurConn = conn
	defer conn.Close()
	transfer := common.Transfer{Conn: conn}
	msmProcess := MsmProcess{}
	for {
		time.Sleep(time.Millisecond * 100)
		msg, err := transfer.ReadPkg()
		if err != io.EOF {
			if err != nil {
				fmt.Println("serverProcessMsg err=", err)
				return
			}
			switch msg.Type {
			case common.NotifyUserStatus:
				updateUserStatus(&msg)
			case common.SMS:
				msmProcess.ReceiveMsg(&msg)
			}
		}
	}
}
