/* 
The Go string is a sequence of variable-width characters.

Go strings and text files occupy less memory or disk space. 
Since, UTF-8 is the standard, Go doesn't need to encode and
decode strings.

Go Strings are value types and immutable. It means if you 
create a string, you cannot modify the contents of the string.
The initial value of a string is empty "" by default.
*/

package main
import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var x string = " Hello, salim  "
	fmt.Println(x) // Hello, salim
	fmt.Println(reflect.TypeOf(x)) // string
	fmt.Println(len(x)) // 12
	fmt.Println("ABC"[1]) // 66 ASCII value of B
	fmt.Println(strings.ToUpper(x))
	fmt.Println(strings.ToLower(x))
	fmt.Println(strings.HasPrefix(x, "Hello")) // true
	fmt.Println(strings.HasSuffix(x, "salim")) // true

	var arr = []string {"It", "is", "good"}
	fmt.Println(strings.Join(arr, " ")) // It is good
	fmt.Println(strings.Repeat(x, 2)) // Hello, salimHello, salim
	fmt.Println(strings.Contains(x, "sa")) // true
	fmt.Println(strings.Index(x, "sa")) // 7
	fmt.Println(strings.Count(x, "l")) // 3
	fmt.Println(strings.Replace(x, "i", "a", 2)) // Hello, salam
    fmt.Println(strings.Split(x, ",")) //[Hello salim]
    fmt.Println(strings.Compare("a", "b")) // -1
    fmt.Println(strings.Compare("a", "a")) // 0
    fmt.Println(strings.Compare("b", "a")) // 1
    fmt.Println(strings.TrimSpace(x)) // Hello, salim
    fmt.Println(strings.ContainsAny(x, "a")) // true
}