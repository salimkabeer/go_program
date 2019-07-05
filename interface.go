/*
Go has different approaches to implement the concepts of object-orientation. 
Go does not have classes and inheritance. Go fulfill these requirements 
through its powerful interface.

Interfaces provide behavior to an object: 
	if something can do this, then it can be used here.

An interface defines a set of abstract methods and does not contain any variable.
syntax:
type Namer interface {  
             Method1(param_list) return_type  
             Method2(param_list) return_type  
             ...  
}  
where Namer is an interface type.

Generally, the name of an interface is formed by the method name plus 
the [e]r suffix, such as Printer, Reader, Writer, Logger, Converter, etc.
	1. A type doesn't have to state explicitly that it implements an interface:
	   interfaces are satisfied implicitly. Multiple types can implement the 
	   same interface.
	2. A type that implements an interface can also have other functions.
	3. A type can implement many interfaces.
	4. An interface type can contain a reference to an instance of any of 
	   the types that implement the interface

*/

package main
import "fmt"

type vehicle interface {
	accelerate()
}

func foo(v vehicle) {
	fmt.Println(v)
}

type car struct {
	model string
	color string
}

func (c car) accelerate() {
	fmt.Println("Accelerating")
}

type toyota struct {
	model string
	color string
	speed int
}

func (t toyota) accelerate() {
	fmt.Println("I am toyota")
}

func main() {
	c1 := car{"suzuki", "blue"}
	t1 := toyota{"Totota", "Red", 100}

	c1.accelerate()
	t1.accelerate()

	foo(c1)
	foo(t1)
}
/*
Output:-
Accelerating
I am toyota
{suzuki blue}
{Totota Red 100}
*/