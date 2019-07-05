/*
In Go, slice is a dynamically-sized, segmented view of an underlying array. 
This segment can be the entire array or a subset of an array. We define the
subset of an array by indicating the start and end index. Slices provide a 
dynamic window onto the underlying array.

odd := [6]int{2,4,6,8,10,12}
var s []int = odd[1:4] // i.e. [4,6,8]

Slice is like reference to an array. Slice does not store any data. 
If we change the elements of an array, it will also modify the 
underlying array. If other slice is referencing the same underlying 
array, their value will also be changed.

*/
package main
import "fmt"

func main() {
	frds := [4]string {"salim", "Akash", "Ruchi", "siri"}
	
	fmt.Println(frds)
	slice1 := frds[0:2]
	slice2 := frds[1:3]
	fmt.Println(slice1,slice2)

	slice2[0] = "xyz"
	fmt.Println(slice1, slice2)
	fmt.Println(frds)
}

/*
Outputs are:-
[salim Akash Ruchi siri]
[salim Akash] [Akash Ruchi]
[salim xyz] [xyz Ruchi]
[salim xyz Ruchi siri]
*/

/*
Slice literal is like an array literal without any length. An example of slice
without length is given below:

s := []struct {
	i int
	b bool
} {{1, true}, {2, true}, {3, false}, {4, true}, {5, false}}

fmt.Println(s) // [{1 true} {2 true} {3 false} {4 true} {5 false}] 

len(slice) // To get length
cap(slice) // To get capacity

Create slice with the help of make function. The make function creates
a zero sized array and returns slice of the array.
	slice := make([]int 10) // length=10 capacity=10 [0 0 0 0 0 0 0 0 0 0]

	
*/