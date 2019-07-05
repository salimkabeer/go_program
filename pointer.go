/*
A pointer is a variable that stores the address of another variable. 
The general form of a pointer variable declaration is:
		var var_name *datatype

A newly declared pointer which has not been assigned to a variable 
has the nil value.

The address-of operator &, when placed before a variable gives us 
the memory address of the variable.

With pointers, we can pass a reference to a variable (for example, 
as a parameter to a function), instead of passing a copy of the 
variable which can reduce memory usage and increase efficiency.

package main
import "fmt"

func main() {
	x := 10
	changeX(&x)
	fmt.Println(x)
}
func changeX(x *int){
	*x = 20
}

output:- 
20
*/

package main
import (
	"fmt"
)

func main() {
	ptr := new(int)
	fmt.Println("Before change ptr", *ptr)
	changePtr(ptr)
	fmt.Println("After change ptr", *ptr)	
}

func changePtr(ptr *int){
	*ptr = 10
}

/*
Output:-
Before change ptr 0
After change ptr 10
*/