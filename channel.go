/*
The channel acts as a pipe by which we send typed values from one Goroutine
to another. It guarantees synchronization since only one Goroutine has access
to a data item at any given time. The ownership of the data is passed between
different Goroutine. Hence, By design it avoids the pitfalls of shared memory
and prevent race condition.

*/

package main  
import "fmt"  
import "time"

func worker(done chan bool) {  
   fmt.Print("working...")  
   time.Sleep(time.Second)  
   fmt.Println("done")  
   done <- true  
}  

func main() {  
   done := make(chan bool, 1)  
   go worker(done)  
   <-done  
}  

/*
Output:-
working...done
*/