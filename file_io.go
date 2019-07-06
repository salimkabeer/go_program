/*
In go os.file objects are used for file manipulations. 
os.File objects are also called filehandles.

open function which is in os package is used to open files in Go. 
ReadFile() in the io/ioutil package is used to read the file. 
This method returns []byte array of read bytes. 
file.WriteString method can be used to write to the file.

We use defer file.close() right after opening the file to make sure
that the file is closed as soon as the function completes. If a file
does not exists or the program has not the sufficient rights to open 
the file then
	inputFile, inputError = os.Open("input.dat") results in an error.

*/

package main
import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Hi...there\n")
	file.WriteString("I am salim\n")
	file.Close()

	stream, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)
	fmt.Println(readString)
}
