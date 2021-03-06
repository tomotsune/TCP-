// @Title  聊天室客户端
// @Description
// @Author  haipinHu  08/10/2021 08:23
// @Update  haipinHu  08/10/2021 08:23
package main

import (
	"awesomeProject/src/client/process"
	"awesomeProject/src/common/model"
	"fmt"
)

func main() {
	key := 0
	loop := true
	for loop {
		fmt.Println("----------欢迎登录多人聊天系统-------------")
		fmt.Println("\t\t\t 1. 登录聊天室")
		fmt.Println("\t\t\t 2. 注册用户")
		fmt.Println("\t\t\t 3. 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("------登录聊天室------")
			// mobile, pwd := "", ""
			member := model.Member{}
			fmt.Println("输入用户的id")
			fmt.Scanln(&member.Mobile)
			fmt.Println("输入用户的密码")
			fmt.Scanln(&member.Pwd)

			userProcess := process.UserProcess{}
			conn, err := userProcess.Login(&member)
			if err != nil {
				fmt.Println("登录失败, err=", err)
			} else {
				go process.ServerProcessMsg(conn)
				process.ShowMenu()
				loop = false
			}
		case 2:
			fmt.Println("--------用户注册-------")
			//mobile, pwd := "", ""
			member := model.Member{}
			fmt.Println("输入用户的id")
			fmt.Scanln(&member.Mobile)
			fmt.Println("输入用户的密码")
			fmt.Scanln(&member.Pwd)

			userProcess := process.UserProcess{}
			err := userProcess.Register(&member)
			if err != nil {
				fmt.Println("注册失败, err=", err)
			}
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入错误")
		}
	}
}
