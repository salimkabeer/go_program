/*
Go Concurrency
Large programs are divided into smaller sub-programs. Programs which run 
their smaller components at the same time is known as concurrency.

Goroutines
The parts of an application that run concurrently are called goroutines. 
Goroutines and channels are used for structuring concurrent programs.

A process is an independently executing entity running in a machine which 
runs in its own address space in memory. A process has threads which are 
simultaneously executing entities. Threads share the same address space of 
the process.

Goroutines are lightweight, much lighter than a thread. Goroutines run in 
the same address space, so access to shared memory must be synchronized; 
This can be done by sync package, but it is recommended to use channels to
synchronize goroutines.

A goroutine is implemented as a function or method. It is called (invoked) 
with the 'go' keyword. When the goroutine finishes, nothing is returned to 
the caller function.

*/

package main
import (
	"fmt"
	"time"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	go fun1()
	go fun2()
	wg.Wait()
}

func fun1() {
	for i:=0; i<10; i++ {
		fmt.Println("fun1 -> ", i)
		time.Sleep(time.Duration(5*time.Millisecond))
	}
	wg.Done()
}

func fun2() {
	for i:=0; i<10; i++ {
		fmt.Println("fun2 -> ", i)
		time.Sleep(time.Duration(10*time.Millisecond))
	}
	wg.Done()
}

/*
Output:-
fun2 ->  0
fun1 ->  0
fun1 ->  1
fun2 ->  1
fun1 ->  2
fun2 ->  2
fun1 ->  3
fun1 ->  4
fun2 ->  3
fun1 ->  5
fun1 ->  6
fun2 ->  4
fun1 ->  7
fun1 ->  8
fun2 ->  5
fun1 ->  9
fun2 ->  6
fun2 ->  7
fun2 ->  8
fun2 ->  9
*/