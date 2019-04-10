package message

//消息类型常量, 每个消息都需要定义一个来一一对应
const (
	SayHelloMessType = "SayHelloMess"
)

//总消息, 最终交互的消息结构体
type Message struct {
	Type string `json:"type"`	//消息的类型
	Content string `json:"content"`	//消息的内容
}

//具体消息, 打招呼
type SayHelloMess struct {
	WhoAmI string `json:"whoAmI"`	//我是谁
	Content string `json:"content"`  //谈话内容
}



