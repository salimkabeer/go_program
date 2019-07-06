/*
Mutual Exclusion locks, or mutexes can be used to synchronize access to state
and safely access data across many goroutines. It acts as a guard to the 
entrance of the critical section of code so that only one thread can enter 
the critical section at a time.

We set a lock around particular lines of code with it. While one Goroutine 
holds the lock, all other Goroutines are prevented from executing any lines 
of code protected by the same mutex, and are forced to wait until the lock 
is yielded before they can proceed.

*/

package main
import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

var wait sync.WaitGroup
var count int
var mutex sync.Mutex

func increment(s string) {
	for i:=0; i<10; i++ {
		mutex.Lock()
		x := count
		x++;
		time.Sleep(time.Duration(rand.Intn(10))*time.Millisecond)
		count = x;
		fmt.Println(s, i, "Count: ", count)
		mutex.Unlock()
	}
	wait.Done()
}

func main() {
	wait.Add(2)
	go increment("foo: ")
	go increment("bar: ")
	wait.Wait()
	fmt.Println("last count value ", count)
}

/*
Output:-
bar:  0 Count:  1
bar:  1 Count:  2
bar:  2 Count:  3
foo:  0 Count:  4
bar:  3 Count:  5
foo:  1 Count:  6
foo:  2 Count:  7
bar:  4 Count:  8
foo:  3 Count:  9
bar:  5 Count:  10
foo:  4 Count:  11
bar:  6 Count:  12
foo:  5 Count:  13
foo:  6 Count:  14
bar:  7 Count:  15
foo:  7 Count:  16
bar:  8 Count:  17
foo:  8 Count:  18
bar:  9 Count:  19
foo:  9 Count:  20
last count value  20
*/
