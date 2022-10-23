package main

import "fmt"

func main() {
	var value string = "This is my text"
	var pointer = &value
	
	fmt.Println(value)
	fmt.Println(*pointer)
}
