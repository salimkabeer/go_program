package main

import "fmt"

func main() {
	// looping in go
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("Print multiple of 2")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i += 2
	}

	x := 5
	for { //infinite loop
		fmt.Println("Infinite")
		x += 3
		if x > 25 {
			break
		}
	}
}
