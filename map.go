/*
In Go, Maps is an unordered collection of key and its associated value. 
They are very good for looking up values fast. The key type must have 
the operations == and != defined, like string, int, float.

Hence arrays, slices and structs cannot be used as key type, 
but pointers and interface types can.

structs can be used as a key when we provide Key() or Hash() method, 
so that a unique numeric or string key can be calculated from the 
struct's fields.

A map is a reference type and declared in general as:
	var map1 map[key-type]value-type
	var map1 map[int]String

*/

package main
import (
	"fmt"
)

func main() {
	x := map[string]int{"abc":24, "xyz":56}
	fmt.Print(x)
	fmt.Println("\n", x["abc"])
	
	// insert into map
	x["ghi"] = 100
	x["hhh"] = 456
	fmt.Println(x)
	
	// update into map element
	x["xyz"] = -40
	x["hhh"] = 300
	fmt.Println(x)

	// delete an element from map
	delete(x, "hhh")
	fmt.Println(x)

    // x["abc"] to retrieve value 
    //  retrive value with bool check
    v1, status1 := x["abc"]
    fmt.Println("value: ", v1, "Status: ", status1)
    v2, status2 := x["jkl"]
    fmt.Println("value: ", v2, "status: ", status2)

}

/*
Output:-
map[abc:24 xyz:56]
 24
map[abc:24 xyz:56 ghi:100 hhh:456]
map[abc:24 xyz:-40 ghi:100 hhh:300]
map[xyz:-40 ghi:100 abc:24]
value:  24 Status:  true
value:  0 status:  false
*/
// ----------------------------------------------------------
/*
map of struct

type vertex struct {
	latitude, longitude float64
}

var m = map[string]Vertex {
	"home": vertex{34.98765, -23.98712},
	"office": vertex{78.12345, 78.65438},
}

func main() {
	fmt.Println(m)
}

*/