/*
we create an anonymous function which acts as a function closure. A function
which has no name is called anonymous function.

A closure is a function which refers reference variable from outside its body.
The function may access and assign to the referenced variables.
*/

package main
import "fmt"

func main() {
	number := 10
	squareNum := func() (int) {
		number *= number
		return number
	}
	fmt.Println(squareNum())
	fmt.Println(squareNum())
}

// Output is:-
// 100
// 10000