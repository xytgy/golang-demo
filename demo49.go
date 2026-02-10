package main

import "fmt"

func main() {
	var i interface{} = "hello"
	if s, ok := i.(string); ok {
		fmt.Println("成功拿到字符串:", s)
	} else {
		fmt.Println("断言失败，i 里面不是字符串")
	}
}
