package main

import "fmt"

func main() {
	x := 15
	a := &x // memory address
	fmt.Println(a)
	fmt.Println(*a)

	*a = 5
	fmt.Println(x)

	*a = *a * *a //multiplication
	fmt.Println(x)
	fmt.Println(*a)
}
