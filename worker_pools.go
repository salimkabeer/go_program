/*
Worker pools is a design in which a fixed number of m workers (Go goroutines) 
works on n tasks in a work queue (Go channel). The work resides in a queue 
until a worker finish its current task and pull a new one.

*/

package main  
import (  
"fmt"  
"time"  
)  
  
func worker(id int, jobs <-chan int, results chan<- int) {  
    for j := range jobs {  
        fmt.Println("worker", id, "processing job", j)  
        time.Sleep(time.Second)  
        results <- j * 2  
    }  
}  
  
func main() {  
    job := make(chan int, 10)  
    result := make(chan int, 10)  
    for w := 1; w <= 2; w++ {  
        go worker(w, job, result)  
    }  
    for j := 1; j <= 9; j++ {  
        job <- j  
    }  
    close(job)  
    for a := 1; a <= 9; a++ {  
        <-result  
    }  
}  

/*
worker 2 processing job 1
worker 1 processing job 2
worker 2 processing job 3
worker 1 processing job 4
worker 2 processing job 5
worker 1 processing job 6
worker 2 processing job 7
worker 1 processing job 8
worker 2 processing job 9

In this example, 2 workers are started and 9 work items are in put onto 
a job channel. Workers have a work loop with a time.Sleep so that each 
ends up working 2 jobs. close is used on the channel after all the work's 
been put onto it, which signals to all 2 workers that they can exit their 
work loop by dropping them out of their loop on range.
*/
