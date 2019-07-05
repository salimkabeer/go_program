/*
In Go programming, calling same function from within the function is known as 
recursion. It is always a good idea to break a problem into multiple tasks.
*/

package main  
import (  
   "fmt"  
)  
func main() {  
   fmt.Println(factorial(5))  
}  
func factorial(num int ) int{  
   if num == 0{  
      return 1  
   }  
   return num*factorial(num-1)  
}  