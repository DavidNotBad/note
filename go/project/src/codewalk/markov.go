package main

//func fibonacci(lst interface{}) func() (interface{}) {
//	num := 0
//	length := len(lst.([] string))
//	return func() (interface{}) {
//		val := lst[num]
//		if num >= length - 1 {
//			num = 0
//		}else{
//			num++
//		}
//		return val
//	}
//}

func fibonacci(lst interface{}) func() {
	switch lst.(type) {
	case []string:
		break
	}

	return nil
}


func main() {
	fibonacci([]string{"1", "2"})

	//lst := []string{"1", "2"}
	//
	//f := fibonacci(lst)
	//for i := 0; i < 10; i++ {
	//	fmt.Println(f())
	//}
}


