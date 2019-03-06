package main

import (
	"chatroom/common"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

//https://www.bilibili.com/video/av35928275/?p=323

//处理登录请求
func serverProcessLogin(conn net.Conn, mes *message.Message)(err error)  {
	//1. 从mes中取出mes.Data， 并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return
	}

	//2. 组装一个返回的消息类型
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//声明一个LoginResMes
	var loginResMes message.LoginResMes
	//判断用户是否合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	}else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在， 请注册再使用。。"
	}
	//组装返回的消息类型
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//3. 发送包
	err = common.WritePkg(conn, data)

	return
}



//根据客户端发送消息种类的不同， 决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message)(err error)  {
	switch mes.Type {
		//处理登录的逻辑
		case message.LoginMesType:
			err = serverProcessLogin(conn, mes)
		//处理注册
		case message.RegisterMesType:

		default:
			fmt.Println("消息类型不存在， 无法处理")
	}
	return
}


//处理和客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()
	fmt.Println("客户端连接成功")

	//读客户端发送的信息
	for {
		msg, err := common.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出， 服务端协程关闭")
			}else {
				fmt.Println("readPkg err=", err)
			}
			return
		}
		fmt.Println("msg=", msg)
		err = serverProcessMes(conn, &msg)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出， 服务端协程关闭")
			}else{
				fmt.Println("serverProcessMes fail", err)
			}
			return
		}
	}
}

func main()  {
	fmt.Println("服务器在8889端口监听。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	//等待客户端连接服务器
	for {
		fmt.Println("等待客户端连接服务器。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		//连接成功后， 启动协程和客户端保持客户端的通讯
		 go process(conn)
	}
}


