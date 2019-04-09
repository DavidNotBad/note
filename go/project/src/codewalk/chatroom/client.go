package main

import (
	"fmt"
	"net"
)






func main() {
	conn, _ := net.Dial("tcp", "0.0.0.0:8812")

	mess := "哈哈水电费"
	n, err := SendPkg(conn, mess)

	fmt.Println(n, err)

	//buffer := make([]byte, 1024)
	//n, _ = conn.Read(buffer)
	//fmt.Println(string(buffer[:n]))
}


