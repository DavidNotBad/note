package main

import "fmt"

type Cat struct {
	Name string
	age int
}

func main() {
	var cat1 Cat
	cat1.Name = "aa"
	cat1.age = 10

	fmt.Println(cat1)
}
