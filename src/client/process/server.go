package process

import (
	"awesomeProject/src/common"
	"fmt"
	"net"
	"os"
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

		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("显示在线用户列表")
			loop = false
		case 4:
			// 不会执行defer
			os.Exit(0)
		default:
			fmt.Println("请选择(1-4)")
		}
	}
}

// 和服务器保持联系
func ServerProcessMsg(conn net.Conn) {
	transfer := common.Transfer{Conn: conn}
	for {
		pkg, err := transfer.ReadPkg()
		if err != nil {
			fmt.Println("serverProcessMsg err=", err)
			return
		}
		fmt.Println(conn.RemoteAddr().String(), "->", pkg)
	}
}
