/*
Go Tickers are used when we want to some work at regular interval of time. 
Tickers can be stopped like timers using the Stop() method.

The NewTicker() method returns a new Ticker having a channel which send the 
time according to the duration argument. The duration must be larger than 
zero, if not, the ticker will panic.

The Tick() is a wrapper for NewTicker which provides access to the ticking 
channel. The Tick() method is useful for clients who don't want to shutdown
the Ticker.

*/

package main  
import "time"  
import "fmt"  

func main() {  
    tickerValue := time.NewTicker(time.Millisecond * 100)  
    go func() {  
        for t := range tickerValue.C {  
            fmt.Println("Tick at", t)  
        }  
    }()  
    time.Sleep(time.Millisecond * 500)  
    tickerValue.Stop()  
    fmt.Println("Ticker stopped")  
}  

/*
Tick at 2019-07-06 10:22:14.047 +0530 IST m=+0.119000001
Tick at 2019-07-06 10:22:14.147 +0530 IST m=+0.219000001
Tick at 2019-07-06 10:22:14.247 +0530 IST m=+0.319000001
Tick at 2019-07-06 10:22:14.347 +0530 IST m=+0.419000001
Ticker stopped
*/
