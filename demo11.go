package main

import "fmt"

// 主函数，程序入口
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

// 交换函数，接收两个字符串参数，返回交	换后的字符串
func swap(x, y string) (string, string) {
	return y, x
}
