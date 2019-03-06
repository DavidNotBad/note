package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your name:")
	fmt.Fscan(inputReader, &key)
	_, err := inputReader.ReadString('\n')
	println(input)
	if err != nil {
		fmt.Println("There")
		return
	}
}