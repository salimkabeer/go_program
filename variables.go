package main

import "fmt"

func add(x float64, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	// Need to specify each return type
	return a, b
}

func main() {
	//var num1, num2 float64 = 5.6, 9.5
	//var num2 float64 = 9.5
	num1, num2 := 5.6, 9.5 // by default uses float64
	fmt.Println(add(num1, num2))

	w1, w2 := "Hey", "there"
	fmt.Println(multiple(w1, w2))

	var a int = 62
	var b float64 = float64(a)
	fmt.Printf("%d \t %f", a, b)
}
