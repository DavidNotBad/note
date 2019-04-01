package process

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

//处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message)(err error)  {
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
	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	loginResMes.Code = 200
	//}else {
	//	loginResMes.Code = 500
	//	loginResMes.Error = "该用户不存在， 请注册再使用。。"
	//}
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在， 请注册后使用..."
	}else{
		loginResMes.Code = 200
		fmt.Println(user, "登录成功")
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
	tf := utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg msg包 failed", err)
	}

	return
}


