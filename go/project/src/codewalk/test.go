package main

import "fmt"

func main() {
	msgCh := make(chan int, 10)
	quitCh := make(chan int, 10)

	msgCh <- 1

	for {
		select {
		case <-msgCh:
			fmt.Println("dowork")
			quitCh <- 2
		case <-quitCh:
			fmt.Println("finish")
			return
		}
	}
}