package main

import "fmt"

func main() {
	// 创建一个容量为2的int类型channel
	ch := make(chan int, 2)
	// 向channel发送值1
	ch <- 1
	// 向channel发送值2
	ch <- 2
	// 关闭channel，表示不再发送数据
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}
