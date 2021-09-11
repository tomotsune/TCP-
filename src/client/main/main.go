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
			mobile, pwd := "", ""
			fmt.Println("输入用户的id")
			fmt.Scanln(&mobile)
			fmt.Println("输入用户的密码")
			fmt.Scanln(&pwd)
			userProcess := process.UserProcess{}
			err := userProcess.Login(&model.Member{Mobile: mobile, Pwd: pwd})
			if err != nil {
				fmt.Println("登录失败, err=", err)
			} else {
				process.ShowMenu()
				loop = false
			}
		case 2:
			fmt.Println("--------用户注册-------")
			mobile, pwd := "", ""
			fmt.Println("输入用户的id")
			fmt.Scanln(&mobile)
			fmt.Println("输入用户的密码")
			fmt.Scanln(&pwd)
			userProcess := process.UserProcess{}
			err := userProcess.Register(&model.Member{Mobile: mobile, Pwd: pwd})
			if err != nil {
				fmt.Println("注册失败, err=", err)
			} else {
				process.ShowMenu()
				loop = false
			}
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入错误")
		}
	}
}
