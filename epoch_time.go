/*
In Go, We can produce the time in seconds, milliseconds, nanosecond. 
The time package has the required methods like secs, nanos etc which 
help us.

The reference time is the unix epoch. We may also convert nanoseconds 
or milliseconds into time format.
*/

package main  
import "fmt"  
import "time"  

func main() {  
  
    p := fmt.Println  
    current_time := time.Now()  
    secs := current_time.Unix()  
    nanos := current_time.UnixNano()  
  
    fmt.Println(current_time)  
  
    millis := nanos / 1000000  
    p(secs)  
    p(millis)  
    p(nanos)  
    p(time.Unix(secs, 0))  
    p(time.Unix(0, nanos))  
  
}  

/*
2019-07-06 10:18:13.224 +0530 IST m=+0.015000001
1562388493
1562388493224
1562388493224000000
2019-07-06 10:18:13 +0530 IST
2019-07-06 10:18:13.224 +0530 IST
*/
