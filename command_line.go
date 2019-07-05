
package main
import "fmt"
import "os"

func main() {
	var str string
	for i:=1; i< len(os.Args); i++ {
		str += os.Args[i] + ""
	}
	fmt.Println(str)
	fmt.Println(os.Args) // returns all args including path
	fmt.Println(os.Args[1:]) // returns all args
}