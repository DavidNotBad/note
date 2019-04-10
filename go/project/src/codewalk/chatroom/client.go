package main

import (
	"fmt"
	"net"
)




func main() {
	//test()
	conn, _ := net.Dial("tcp", "0.0.0.0:8812")

	message := Message{}
	data, err := message.Marshal(SayHelloMessType, SayHelloMess{
		WhoAmI: "尼尔",
		Content: "hello world!",
	})
	n, err := SendPkg(conn, data)

	fmt.Println(n, err)
}


