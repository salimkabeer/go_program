/* It defines the pkg name in which program should lie, Its a mandatory stmt.
As Go run in pkg, the main pkg is the starting point to run the program. 
Each pkg has a path and name associated with it*/

package main

import "fmt" //It's a preprocessor command to include pkg "fmt"

func main() { //program execution start from here
    fmt.Println("Hello, World!")
}
