package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {

}

//请求注册
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//1. 连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	//2. 准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3. 创建一个RegisterMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4. 将loginMes序列化
	data, err := json.Marshal(registerMes)
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

	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息出错了 err=", err)
	}


	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功")
	}else{
		fmt.Println(registerResMes.Error)
	}

	os.Exit(0)
	return
}



//写一个函数， 完成登录
func (this *UserProcess) Login(userId int, userPwd string)(err error) {
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
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")

		//显示当前在线用户列表， 遍历loginResMes.UserId
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UsersId {
			fmt.Println("用户id:\t", v)
		}
		fmt.Printf("\n\n")

		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务器端的通讯， 如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端
		go serverProcessMes(conn)

		//1. 显示我们的登录成功的菜单
		for {
			ShowMenu()
		}
	}else{
		fmt.Println(loginResMes.Error)
	}

	return
}