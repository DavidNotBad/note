package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type server struct {
	network string			//网络连接类型
	address string			//地址+端口
}

//开启服务器监听, 获取客户端连接, 接收客户端发送过来的数据
func (server *server) listen()(err error) {
	//建立服务器进行端口监听
	listener, err := net.Listen(server.network, server.address)
	if err != nil {
		return
	}

	//延时关闭服务器
	defer func() {
		err = listener.Close()
	}()

	for {
		//等待客户端连接上来, 这里会阻塞
		conn, err := listener.Accept()
		if err != nil {
			break
		}

		//后面加上协程process
		//读取包, 如果出现错误, 则丢包
		mess, err := ReadPkg(conn)
		if err != nil {
			continue
		}

		message := Message{}
		err = message.UnMarshal(mess)

		//业务处理, 调用各种process, 如userProcess, 使用switch语句??, 使用route??, message也可以使用route?
		err = server.accept(conn, message)
	}

	return
}




//客户端发送消息过来, 进行业务处理
func (server *server) accept(conn net.Conn, message Message)(err error) {
	fmt.Println("收到客户端", conn.RemoteAddr(),"发送的消息", message)
	fmt.Println(message.Content)

	var sayHelloMess SayHelloMess
	err = json.Unmarshal([]byte(message.Content), &sayHelloMess)
	if err != nil {
		return
	}

	//messageContent, _ := message.UnMarshalContent()
	//messageContent.(SayHelloMess)

	//fmt.Println(mess.WhoAmI)
	//fmt.Println(mess.Content)
	return
}



func main()  {
	//sayHelloMess := &SayHelloMess{}
	//mess := "{\"whoAmI\":\"尼尔\",\"content\":\"hello world!\"}"
	//err := json.Unmarshal([]byte(mess), sayHelloMess)
	//fmt.Println(err)
	//os.Exit(0)

	//监听端口
	server := server{
		"tcp", "0.0.0.0:8812",
	}
	err := server.listen()
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
}
