package main

import (
	"chatroom/common"
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//写一个函数， 完成登录
func login(userId int, userPwd string)(err error) {
	//开始定协议
	//fmt.Printf("userId = %d userPwd=%s\n", userId, userPwd)
	//return nil

	//1. 连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	//2. 准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3. 创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4. 将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5. 把data赋给mes.Data字段
	mes.Data = string(data)

	//6. 将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//7. 发送通讯内容的长度给服务器， 用来验证是否丢包
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], uint32(len(data)))
	n, err := conn.Write(buf[:])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	fmt.Println("客户度发送消息的长度", len(data), "ok")

	//8. data就是可以发送的消息
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	fmt.Println("客户度发送消息", string(data), "ok")

	//9. 还需要处理服务器端返回的消息
	mes, err = common.ReadPkg(conn)
	if err != nil {
		fmt.Println("readPkg err", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	}else{
		fmt.Println(loginResMes.Error)
	}

	return
}


