package main

import "fmt"

func main() {
	// 1) loop limit 1
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 2) loop limit 2
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	// 3) while loop 1 
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// 4) while loop 2 
	flag := true 
	sum = 0
	for flag {
		if sum > 1000 {
			flag = false
		}
		sum++
	}
	fmt.Println(sum)


	// 4) loop forever
	sum = 0
	for {
		if sum > 1000 {
			fmt.Println("while break")
			break
		}
		sum++
	}
}

