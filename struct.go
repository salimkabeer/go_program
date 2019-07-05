/*
In Go, Struct can be used to create user-defined types.

Struct is a composite type means it can have different properties 
and each property can have their own type and value.

Struct can represent real-world entity with these properties. We 
can access a property data as a single entity. It is also valued
types and can be constructed with the new() function.

package main  
import (  
   "fmt"  
)  
type person struct {  
   firstName string  
   lastName  string  
   age       int  
}  
func main() {  
   x := person{age: 30, firstName: "John", lastName: "Anderson", }  
   fmt.Println(x)  
   fmt.Println(x.firstName)  
}  

Output:-
{John Anderson 30}
John

Struct is a data type and can be used as an anonymous field (having 
only the type). One struct can be inserted or "embedded" into other 
struct.

It is a simple 'inheritance' which can be used to implement impleme-
ntations from other type or types.
*/

package main
import "fmt"

type person struct {
	fname string
	lname string
}

type employee struct {
	person
	empid int
}

func (p person) details() {
	fmt.Println(p, " " + "I am a person")
}

func (e employee) details() {
	fmt.Println(e, " " +  "I am an employee")
}

func main() {
	p1 := person{"salim", "kabeer"}
	p1.details()

	e1 := employee{person:person{"John", "Mclane"}, empid:10}
	e1.details()
}
/*
Output:-
{salim kabeer}  I am a person
{{John Mclane} 10}  I am an employee
*/