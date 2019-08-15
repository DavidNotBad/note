package main

import (
	"chatroom/common/message"
	process2 "chatroom/server/process"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//根据客户端发送消息种类的不同， 决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	//处理登录的逻辑
	case message.LoginMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	//处理注册
	case message.RegisterMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	default:
		fmt.Println("消息类型不存在， 无法处理")
	}
	return
}

func (this *Processor) process() (err error) {
	//读客户端发送的信息
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	for {
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出， 服务端协程关闭")
			} else {
				fmt.Println("readPkg err=", err)
			}
			return err
		}
		fmt.Println("msg=", msg)
		err = this.serverProcessMes(&msg)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出， 服务端协程关闭")
			} else {
				fmt.Println("serverProcessMes fail", err)
			}
			return err
		}
	}
}
