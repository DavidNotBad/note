package main

import (
	"bufio"
	"fmt"
	"os"
)

var userId int
var userPwd string

func main()  {
	//接收用户的输入
	var key int
	//判断是否还继续显示菜单
	var loop = true
	stdin := bufio.NewReader(os.Stdin)

	for loop {
		fmt.Println("----------欢迎登陆多人聊天系统------")
		fmt.Println("1. 登陆聊天室")
		fmt.Println("2. 注册用户")
		fmt.Println("3. 退出系统")
		fmt.Println("请选择： ")

		fmt.Fscan(stdin, &key)
		stdin.ReadString('\n')

		loop = false
		switch key {
		case 1:
			fmt.Println("1. 登陆聊天室")
		case 2:
			fmt.Println("2. 注册用户")
		case 3:
			fmt.Println("退出系统成功")
			os.Exit(0)
		default:
			loop = true
			fmt.Println("请输入1-3的值")
			continue
		}
	}

	//根据用户的输入， 显示新的信息
	if key == 1 {
		//说明用户要登录
		fmt.Println("请输入用户的id")
		fmt.Fscan(stdin, &userId)
		stdin.ReadString('\n')
		fmt.Println("请输入用户的密码")
		fmt.Fscan(stdin, &userPwd)
		stdin.ReadString('\n')
		//先把登录的函数， 写到另外一个文件
		login(userId, userPwd)
	}else if key == 2{
		fmt.Println("注册用户")
	}

}