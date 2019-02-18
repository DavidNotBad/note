package main

import "fmt"

func main() {
	var cities = make(map[string] string)
	cities["a"] = "aa"
	cities["b"] = "bb"

	for val, key := range cities {
		fmt.Println(val)
		fmt.Println(key)
	}
}
