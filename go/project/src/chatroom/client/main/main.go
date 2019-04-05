package main

import (
	"bufio"
	"chatroom/client/process"
	"fmt"
	"os"
)

var userId int
var userPwd string
var userName string

func main()  {
	//接收用户的输入
	var key int
	stdin := bufio.NewReader(os.Stdin)

	for true {
		fmt.Println("----------欢迎登陆多人聊天系统------")
		fmt.Println("1. 登陆聊天室")
		fmt.Println("2. 注册用户")
		fmt.Println("3. 退出系统")
		fmt.Println("请选择： ")

		fmt.Fscan(stdin, &key)
		stdin.ReadString('\n')

		switch key {
		case 1:
			fmt.Println("1. 登陆聊天室")//说明用户要登录
			fmt.Println("请输入用户的id")
			fmt.Fscan(stdin, &userId)
			stdin.ReadString('\n')
			fmt.Println("请输入用户的密码")
			fmt.Fscan(stdin, &userPwd)
			stdin.ReadString('\n')

			//先把登录的函数， 写到另外一个文件
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("2. 注册用户")

			fmt.Println("请输入用户的id")
			fmt.Fscan(stdin, &userId)
			stdin.ReadString('\n')

			fmt.Println("请输入用户的密码")
			fmt.Fscan(stdin, &userPwd)
			stdin.ReadString('\n')

			fmt.Println("请输入用户的名字")
			fmt.Fscan(stdin, &userName)
			stdin.ReadString('\n')

			//注册
		case 3:
			fmt.Println("退出系统成功")
			os.Exit(0)
		default:
			fmt.Println("请输入1-3的值")
			continue
		}
	}


}