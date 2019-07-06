/*
Go has very good support for the time manipulations. Unix epoch 
time is used as reference for time manipulations.

we can built a time object by using Date method provided in the 
time package. The package contains methods like year(), month(), 
day(), location() etc.

We invoke these methods by using time object.
*/

package main  
import "fmt"  
import "time"  
  
func main() {  
    p := fmt.Println  
  
    present := time.Now()// current time  
    p(present)  
  
    DOB := time.Date(1993, 02, 28, 9,04,39,213 ,time.Local)  
    fmt.Println(DOB)  
  
    fmt.Println(DOB.Year())  
    fmt.Println(DOB.Month())  
    fmt.Println(DOB.Day())  
    fmt.Println(DOB.Hour())  
    fmt.Println(DOB.Minute())  
    fmt.Println(DOB.Second())  
    fmt.Println(DOB.Nanosecond())  
    fmt.Println(DOB.Location())  
  
    fmt.Println(DOB.Weekday())  
  
    fmt.Println(DOB.Before(present))  
    fmt.Println(DOB.After(present))  
    fmt.Println(DOB.Equal(present))  
  
    diff := present.Sub(DOB)  
    fmt.Println(diff)  
    fmt.Println(diff.Hours())  
    fmt.Println(diff.Minutes())  
    fmt.Println(diff.Seconds())  
    fmt.Println(diff.Nanoseconds())  
    fmt.Println(DOB.Add(diff))  
    fmt.Println(DOB.Add(-diff))  
}  

/*
2019-07-06 10:11:32.93 +0530 IST m=+0.014000001
1993-02-28 09:04:39.000000213 +0530 IST
1993
February
28
9
4
39
213
Local
Sunday
true
false
false
230977h6m53.929999787s
230977.1149805555
1.3858626898833329e+07
8.315176139299998e+08
831517613929999787
2019-07-06 10:11:32.93 +0530 IST
1966-10-24 07:57:45.070000426 +0530 IST
*/