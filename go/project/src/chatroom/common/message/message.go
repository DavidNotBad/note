package message

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)


type Message struct {
	Type string `json:"type"` //消息的类型
	Data string `json:"data"` //消息的类型
}

//登录消息
type LoginMes struct {
	UserId int `json:"userId"` //用户id
	UserPwd string `json:"userPwd"` //用户密码
	UserName string `json:"userName"` //用户名
}

//登录返回状态
type LoginResMes struct {
	Code int `json:"code"` //返回状态码500：为注册，200：登录成功
	Error string `json:"error"` //返回错误信息
}

//注册
type RegisterMes struct {

}