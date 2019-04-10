package main

import (
	"encoding/json"
)


//总消息, 最终交互的消息结构体
type Message struct {
	Type string `json:"type"`	//消息的类型
	Content string `json:"content"`	//消息的内容
}

//将消息序列化为切片
func (message *Message) Marshal(messType string, messageContent MessageContent)(mess []byte, err error)  {
	//序列化消息的内容
	data, err := json.Marshal(messageContent)
	if err != nil {
		return
	}

	//拼接并序列化总消息
	message.Type = messType
	message.Content = string(data)

	//序列化总消息
	mess, err = json.Marshal(message)
	return
}

//将消息序列化为切片
func (message *Message) UnMarshal(data []byte)(err error)  {
	//反序列化消息的内容
	//这里使用了指针, 会直接改变该结构体, 所以message.Type和message.Content会被赋值
	err = json.Unmarshal(data, message)
	if err != nil {
		return
	}

	return
}






//具体消息, 序列化后放到总消息的Data字段
//暂定为空接口, 以后根据业务需要添加接口方法
type MessageContent interface {
}


//消息类型常量, 每个消息都需要定义一个来一一对应
const (
	SayHelloMessType = "SayHelloMess"
)



//具体消息, 打招呼
type SayHelloMess struct {
	WhoAmI string `json:"whoAmI"`	//我是谁
	Content string `json:"content"`  //谈话内容
}

