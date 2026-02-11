package main

import "fmt"

func main() {
	var i interface{} = 42
	if str, ok := i.(string); ok {
		fmt.Println("String:", str)
	} else {
		fmt.Println("Not a string")
	}
}
