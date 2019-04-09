package main

import (
	"errors"
)


//具体消息, 序列化后放到总消息的Data字段
//暂定为空接口, 以后根据业务需要添加接口方法
type MessageContent interface {
}


//消息类型常量, 每个消息都需要定义一个来一一对应
const (
	SayHelloMessType = "SayHelloMess"
)

//根据消息type, 获取具体的消息实例
func getInstance(insType string) (messCont MessageContent, err error) {
	switch insType {
		case SayHelloMessType:
			messCont = SayHelloMess{}
		default:
			err = errors.New("没有找到相应的实例")
	}
	return
}


//具体消息, 打招呼
type SayHelloMess struct {
	WhoAmI string `json:"whoAmI"`	//我是谁
	Content string `json:"content"`  //谈话内容
}


