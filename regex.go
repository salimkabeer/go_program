/*
Go Regex package is used for searching string. To search a string, 
we need to provide a pattern of the string.

We need to compile the pattern into the regex object so that we can
invoke methods through it.

The regex object can be retrieved by using compile() and mustcompile() 
functions. Now we can use functions to find strings such as FindString(), 
FindStringSubmatch(), FindStringIndex() etc.

*/

package main
import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(".com")

	fmt.Println(re.FindString("google.com")) // .com
	fmt.Println(re.FindString("abc.org"))  // empty
	fmt.Println(re.FindString("hi.com"))  // .com

    fmt.Println(re.FindStringIndex("google.com")) // [6 10]
    fmt.Println(re.FindStringIndex("abx.org"))  // []

    re1 := regexp.MustCompile("f([a-z]+)ing")
    fmt.Println(re1.FindStringSubmatch("hahaflying")) // [flying ly]
    fmt.Println(re1.FindStringSubmatch("abcsdfloatinganga")) // [floating loat]
}
// The FindString() method returns a string having the text of the left 
//most match. If no match is found, empty string is returned.