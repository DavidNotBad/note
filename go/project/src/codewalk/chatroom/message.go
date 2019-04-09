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
func (message *Message) UnMarshal(data []byte)(messageContent MessageContent, err error)  {
	//反序列化消息的内容
	//这里使用了指针, 会直接改变该结构体, 所以message.Type和message.Content会被赋值
	err = json.Unmarshal(data, message)
	if err != nil {
		return
	}

	//动态获取消息实例
	instance, err := getInstance(message.Type)

	//反序列化具体消息
	err = json.Unmarshal([]byte(message.Content), instance)
	return
}






