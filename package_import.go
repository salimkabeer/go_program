/*
Packages are used to categorize our program so that it can be easy to maintain.
Every go-file belongs to some package. Each Go application must have "main" 
package so that it can be compiled.

An application can consist of different packages. Many different .go file can
belong to one main package.

We can save Go program by any name but it must have main package. The package
name should be written in lowercase letters.

If a package is changed and recompiled, all the client programs that use this
package must be recompiled too!

-------------------------------------------------------------------------------

A Go program is linked to different packages through the import keyword.

The package names are enclosed within double quotes "". Import loads the public
declarations from the compiled package, it does not insert the source code.

We can import multiple packages by a separate statement like:
Example:-
	import "fmt"
	import "os"
		OR
	import "fmt"; import "os"
		OR
	import (
		"fmt"
		"os"
	)
		OR
	import ( "fmt"; "os")

-------------------------------------------------------------------------------

Visibility
An identifier can be variable, constant, function, type or struct field. We can
declare identifier in lowercase or uppercase letters.

If we declare identifier in lowercase letter, it will be visible within the 
package only. But if we declare package in uppercase letter, it will be visible
within and outside the package which is also known as exported.

The dot . Operator is used to access the identifier e.g. pack.Age where pack 
is the package name and Age is the identifier.	
*/

package main

import ( "log"; "os" )

var (
    newFile *os.File
    err     error
)

func main() {
    newFile, err = os.Create("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    log.Println(newFile)
    newFile.Close()
}