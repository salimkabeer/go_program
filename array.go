/*
In Go, an array is a homogeneous data structure (Fix type) and has a 
fixed-length. The type can be anything like integers, string or self-
defined type.

The items in the array can be accessed through their index, It starts with zero.
The number of items in the array is called the length or size of an array. It is
fixed and must be declared in the declaration of an array variable.

var <identifier> [len]type
var arr_name [10]int

*/

package main
import "fmt"

func main() {
	var x [5]int
	var i, j int
	for i=0; i<5; i++ {
		x[i] = i + 10
	}
	for j=0; j<5; j++ {
		fmt.Printf("Index %d has %d\n", j, x[j])
	}
}