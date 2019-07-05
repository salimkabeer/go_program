/*
Go Reflection is the ability of a program to examine its own structure, 
particularly through the types; it's a form of meta-programming.

Reflect can be used to investigate types and variables at runtime, 
e.g. its size, its methods, and it can also call these methods 'dynamically'.


*/
package main
import (
	"fmt"
	"reflect"
)

func main() {
	age := 27.5
	fmt.Printf("%T\n", age)
	fmt.Println(reflect.TypeOf(age))
}

/*
Output:-
float64
float64
*/