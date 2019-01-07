package main

import "fmt"

func main() {
    var x float64 = 20.0
    y := 44
    fmt.Printf("Type of x : %T\n", x)
    fmt.Printf("Type of y : %T\n", y)

    var a, b = 4, "foo"
    fmt.Printf("Type of a : %T\n", a)
    fmt.Printf("Type of b : %T\n", b)
}
