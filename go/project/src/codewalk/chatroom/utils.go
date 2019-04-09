package main

import (
	"errors"
	"net"
	"strconv"
)


// 接收第三方发送的消息
// 包最大的长度为4位数, 如果包读取失败, 则clientMes为空
// 发送的包的规则为: 1. 先发送包长度(字符串), 2. 再发送包内容(字符串)
func ReadPkg(conn net.Conn) (clientMess string, err error) {
	//1: 先获取第三方发送的数据包长度
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	//设定数据包长度为四位数, 即最大9999个长度, 超出则报错
	//这个位数4可以考虑放到配置里
	if n > 4 {
		err = errors.New("数据包的长度超出限制")
		return
	}

	//获取数据包长度(不是数据包长度的位数)
	pkgLen, err := strconv.Atoi(string(buffer[:n]))
	//设置读取包的次数
	readCount := pkgLen / 1024 + 1

	//循环读取客户端提交的数据, 每次读取1024字节, 直到读取完毕为止
	for i := 0; i < readCount; i++ {
		//如果客户端没有断掉连接, 则会阻塞
		//注意: 读取完毕后, 如果客户端主动断掉连接, 则不能向客户端写数据
		n, err := conn.Read(buffer)

		//读取失败, 不再读取
		if n <= 0 {
			err = errors.New("读取失败")
			break
		}
		if err != nil {
			break
		}

		clientMess += string(buffer[:n])
	}

	//出错则清空字符串
	if err != nil {
		clientMess = ""
	}
	return
}


//发送数据包到第三方
//这里没有对数据包的长度进行限制, 但是接收方需要限制, 可以考虑把长度的限制加上
//在这里做了规范, 发送的包在转成 []byte 类型之前都需要转成字符串(重要)
//1. 先发送数据包的长度
//2. 再发送数据包
func SendPkg(conn net.Conn, message string)(n int, err error)  {
	//先发送数据包的长度(字符串)
	_, err = conn.Write([]byte(strconv.Itoa(len(message))))

	//再发送数据包的内容
	n, err = conn.Write([]byte(message))

	return
}

