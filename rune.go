/*
The type rune is an alias for type int32. 
The rune literals are an integer value.

If you store string value in rune literal, 
it will provide the ASCII value of the character. 
For example, the rune literal of 'A' will be 65.

*/

package main
import (
	"fmt"
	"reflect"
)

func main() {
	rune := 'A'
	fmt.Printf("%d \n", rune)
	fmt.Println(reflect.TypeOf(rune))
}

/*
Output:-
65
int32
*/