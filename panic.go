/*
Go panic is a mechanism by which we handle error situations. 
Panic can be used to abort a function execution. When a function
calls panic, its execution stops and the control flows to the 
associated deferred function.

The caller of this function also gets terminated and caller's 
deferred function gets executed (if present any). This process
continues till the program terminates. Now the error condition 
is reported.

This termination sequences is called panicking and can be 
controlled by the built-in function recover.

package main  
  
import "os"  
  
func main() {  
    panic("Error Situation")  
    _, err := os.Open("/tmp/file")  
    if err != nil {  
        panic(err)  
    }  
}  

Output:-
panic: Error Situation

goroutine 1 [running]:
main.main()
/Users/pro/GoglandProjects/Panic/panic example1.go:6 +0x39

*/

package main
import (
	"fmt"
)

func main() {
	fmt.Println("Calling x from main")
	x()
	fmt.Println("Returned from x")
}

func x() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering in x", r)
		}
	}()
	fmt.Println("Executing x ...")
	fmt.Println("Calling y from x")
	y(0)
	fmt.Println("Returned from y")
}

func y(i int) {
	fmt.Println("Executing y ...")
	if i > 2 {
		fmt.Println("panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in y", i)
	fmt.Println("Printing in y", i)
	y(i + 1)
}

/*
Output:-
Calling x from main
Executing x ...
Calling y from x
Executing y ...
Printing in y 0
Executing y ...
Printing in y 1
Executing y ...
Printing in y 2
Executing y ...
panicking
Defer in y 2
Defer in y 1
Defer in y 0
Recovering in x 3
Returned from x
*/