package main

import (
	"fmt"
)

func main() {

	//第一种，指定变量类型，如果没有初始化，则变量默认为零值。
	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Printf("i: %d, f: %f, b: %v, s: %s\n", i, f, b, s)

	//第二种，根据值自行判定变量类型。
	// var d = true
	// fmt.Println(d)

	//第三种，如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误
	f := "Runoob" // var f string = "Runoob"
	fmt.Println(f)
}
