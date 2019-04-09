package main

import (
	"fmt"
	"io"
	"net"
	"os"
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

		//后面加上协程
		//读取包
		clientMess, err := ReadPkg(conn)//修改. 返回切片
		//序列化包, 解析内容WhoAmI  Content
		fmt.Println(clientMess)

		os.Exit(0)
		//业务处理
		err = server.accept(conn, clientMess)
	}

	return
}




//客户端发送消息过来, 进行业务处理
func (server *server) accept(conn net.Conn, mess string)(err error) {
	fmt.Println("收到客户端", conn.RemoteAddr(),"发送的消息", mess)
	n, err := conn.Write([]byte("你好啊"))
	fmt.Println(err == io.EOF)
	fmt.Println(n, err)
	return
}



func main()  {
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
