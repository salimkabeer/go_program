package main

import (
	"fmt"
	"math"
	"math/rand"
)

func foo() {
	fmt.Println("The square root of 4 is ", math.Sqrt(4))
}

func main() {
	foo()
	fmt.Println("A number between 1-100 is :", rand.Intn(100))
}
