package main

import (
	"codewalk/frame/common"
	"codewalk/frame/common/message"
	"fmt"
	"net"
)



//go build -o client.exe codewalk/frame/client/main
func main() {
	//test()
	conn, _ := net.Dial("tcp", "0.0.0.0:8812")

	data, err := common.MessageMarshal(message.SayHelloMessType, message.SayHelloMess{
		WhoAmI: "尼尔",
		Content: "hello world!",
	})
	n, err := common.SendPkg(conn, data)

	fmt.Println(n, err)
}


