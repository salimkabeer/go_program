/*
var arrayname [n][m]type

a = [3][4]int

a = [2][3]int {
	{2.4,6},
	{10,20,30}
}

int x = a[1][2]

*/

package main 
import "fmt"

func main() {
	var a = [3][3] int {{1,2,3},{4,5,6},{7,8,9}}
	var i, j int

	for i=0; i<3; i++ {
		for j=0; j<3; j++ {
			// print transpose of a matrix
			fmt.Printf("\t%d", a[j][i])
		}
		fmt.Println("")
	}
}