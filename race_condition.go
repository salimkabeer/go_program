/*
A race condition occurs in Go when two or more goroutines try to access the 
same resource. It may happen when a variable attempts to read and write the 
resource without any regard to other routines.

*/

package main
import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

var wait sync.WaitGroup
var count int

func increment(s string) {
	for i:=0; i<10; i++ {
		x := count
		x++;
		time.Sleep(time.Duration(rand.Intn(4))*time.Millisecond)
		count = x;
		fmt.Println(s, i, "Count: ", count)
	}
	wait.Done()
}

func main() {
	wait.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	wait.Wait()
	fmt.Println("last count value", count)
}

/*
Output:-
bar:  0 Count:  1
foo:  0 Count:  1
bar:  1 Count:  2
foo:  1 Count:  2
bar:  2 Count:  3
foo:  2 Count:  4
foo:  3 Count:  6
foo:  4 Count:  7
foo:  5 Count:  8
bar:  3 Count:  5
foo:  6 Count:  9
bar:  4 Count:  10
foo:  7 Count:  11
foo:  8 Count:  13
bar:  5 Count:  12
foo:  9 Count:  14
bar:  6 Count:  15
bar:  7 Count:  16
bar:  8 Count:  17
bar:  9 Count:  18
last count value 18
*/
