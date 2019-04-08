package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8812")
	n, err := conn.Write([]byte("123456789"))
	fmt.Println(n, err)
}


