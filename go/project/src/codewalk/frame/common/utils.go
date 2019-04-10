package common

import (
	"bytes"
	"codewalk/frame/common/message"
	"encoding/json"
	"net"
	"strconv"
)


// 接收第三方发送的消息
// 包最大的长度为4位数, 如果包读取失败, 则clientMes为空
// 发送的包的规则为: 1. 先发送包长度(字符串), 2. 再发送包内容(字符串)
func ReadPkg(conn net.Conn) (mess []byte, err error) {
	//1: 先获取第三方发送的数据包长度
	//数据包长度已经和客户端约定好是8位
	bufferLen := make([]byte, 8)
	_, err = conn.Read(bufferLen)
	if err != nil {
		return
	}

	//获取数据包长度(不是数据包长度的位数)
	pkgLen, err := strconv.Atoi(string(bytes.Trim(bufferLen[:8], "\x00")))
	if err != nil {
		return
	}

	//设置读取包的次数
	readCount := pkgLen / 1024 + 1

	//循环读取客户端提交的数据, 每次读取1024字节, 直到读取完毕为止

	buffer := make([]byte, 1024)
	for i := 0; i < readCount; i++ {
		//如果客户端没有断掉连接, 则会阻塞
		//注意: 读取完毕后, 如果客户端主动断掉连接, 则不能向客户端写数据
		n, err := conn.Read(buffer)

		//读取失败, 不再读取
		if n <= 0 {
			break
		}
		if err != nil {
			break
		}

		mess = append(mess, buffer[:n]...)
	}

	//出错则清空字符串
	if err != nil {
		mess = []byte(nil)
	}
	return
}


//发送数据包到第三方
//1. 先发送数据包的长度, 长度固定占8位
//2. 再发送数据包
func SendPkg(conn net.Conn, mess []byte)(n int, err error)  {
	//1. 先发送数据包的长度
	//这里做了规范, 前面8位是长度, 不够补零
	var pkgLen = make([]byte, 8)
	pkgLen = []byte(strconv.Itoa(len(mess)))
	_, err = conn.Write(pkgLen[:8])

	//2. 再发送数据包的内容
	n, err = conn.Write(mess)

	return
}


//组装Message, 序列化为切片
func MessageMarshal(messType string, messageContent interface{})(mess []byte, err error)  {
	//序列化消息的内容
	data, err := json.Marshal(messageContent)
	if err != nil {
		return
	}

	msg := message.Message{}
	//拼接并序列化总消息
	msg.Type = messType
	msg.Content = string(data)

	//序列化总消息
	mess, err = json.Marshal(&msg)
	return
}




