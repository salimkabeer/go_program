/*
In Go, functions are the basic building blocks. A function is used to break a 
large problem into smaller tasks. We can invoke a function several times, hence
functions promote code reusability. There are 3 types of functions in Go:
	Normal functions with an identifier
	Anonymous or lambda functions
	Method (A function with a receiver)

Function cannot be declared inside another function. If we want to achieve 
this, we can do this by anonymous function.
*/

// Go Function Example
package main
import "fmt"

type Employee struct {
	fname string
	lname string
}

func (emp Employee) fullname() {
	fmt.Println(emp.fname+" "+emp.lname)
}

func main() {
	e1 := Employee{"salim", "kabeer"}
	e1.fullname()

    fmt.Println(fun1)
    fmt.Println("-------------------")
	fmt.Println(add_sub_all(10,15,-40))
}

//Go function with return
func fun1() int {
	return 123456789
}

//Go function with multiple return
func add_sub_all(args ... int) (int, int) {
	add := 0
	sub := 0
	for _, value := range args {
		add += value
		sub -= value
	}
	return add, sub
}

/*
Outputs are:-
salim kabeer
0x48e910
-------------------
-15 15
*/