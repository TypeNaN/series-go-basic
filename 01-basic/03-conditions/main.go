package main

import "fmt"

func main() {
	var myBool bool = false
	var myNumber int = 2

	if myBool {
		fmt.Println("myBool == true")
	} else if myNumber == 4{
		fmt.Println("myNumber == 4")
	} else {
		fmt.Println("print here on else")
	}
}
