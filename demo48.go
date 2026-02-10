package main

import "fmt"

func printValue(val interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", val, val)
}

func main() {
	printValue(42)           // int
	printValue("hello")      // string
	printValue(3.14)         // float64
	printValue([]int{1, 2})  // slice
	printValue([2]int{1, 2}) // array
}
