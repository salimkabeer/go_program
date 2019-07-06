package main  
  
import "fmt"  
import (  
    "math/rand"  
    "time"  
)  

func main() {  
    // will produce a random integer between 0 to 100
    fmt.Print(rand.Intn(100))  
    fmt.Println()  
  
    // will produce a random number between 0 to 1  
    fmt.Print(rand.Float64())
    fmt.Println()  
      
    // seeding do that a random number will produced  
    rand.Seed(time.Now().Unix())
    myrand := random(1, 20)  
  
    fmt.Println(myrand)  
  
}  
  
func random(min, max int) int {  
    return rand.Intn(max - min) + min  
}  