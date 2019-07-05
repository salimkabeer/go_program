/*
Go does not have an exception mechanism like try/catch in Java, 
we cannot throw an exception in Go.

Go uses a different mechanism which is known as 
defer-panic-and-recover mechanism.

Go handles simple errors for function, methods by returning an 
error object. The error object may be the only or the last return
value. The error object is nil if there is no error in the function.

We should always check the error at the calling statement, 
if we receive any of it or not.

We should never ignore errors, it may lead to program crashes.

The way go detect and report the error condition is

A function which can result in an error returns two variables: a value
and an error-code which is nil in case of success, and != nil in case 
of an error-condition. The error is checked, after the function call. 
In case of an error ( if error != nil), the execution of the actual 
function (or if necessary the entire program) is stopped.
Go has predefined error interface type

	type error interface {
		Error() string
	}

	err := errors.New("math square root of negative number")
*/

package main
import (
	"fmt"
	"errors"
	"math"
)

func Sqrt(value float64) (float64, error) {
	if (value < 0) {
		return 0, errors.New("Math: negative number")
	}
	return math.Sqrt(value), nil
}

func main() {
	result, err := Sqrt(-64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = Sqrt(64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

/*
Output:-
Math: negative number
8
*/