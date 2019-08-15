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
	Conn   net.Conn
	UserId int
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1. 从mes中取出mes.Data， 并直接反序列化成RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return
	}

	//2. 组装一个返回的消息类型
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//声明一个RegisterResMes
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		fmt.Println("register err = ", err)
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}

	//反序列化返回的消息类型
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}
	//将data服务值resMes
	resMes.Data = string(data)

	//对resMes进行序列化， 准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail, err=", err)
		return
	}

	//发送data， 我们将其封装到writePkg函数
	//因为使用分成模式（mvc）， 我们先创建一个Transfer实例， 然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	return
}

//处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
	} else {
		loginResMes.Code = 200

		//把登录成功的用户放入到userMgr中
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		for id := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}

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
