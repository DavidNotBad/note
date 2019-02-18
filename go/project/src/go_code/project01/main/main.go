package main

import "fmt"

func main() {
	//创建一个map切片
	var mapSlice = make([]map[string]string, 2)
	//第一种方式给map切片的第一个元素赋值
	mapSlice[0] = make(map[string]string, 2)
	mapSlice[0]["name"] = "a"
	mapSlice[0]["age"] = "10"
	//第二种方式给map切片的第二个元素赋值
	mapSlice[1] = map[string]string{
		"name": "b",
		"age":  "20",
	}
	//切片追加元素
	mapSlice = append(mapSlice, map[string]string{
		"name": "c",
		"age":  "30",
	})
	fmt.Println(mapSlice)
}
