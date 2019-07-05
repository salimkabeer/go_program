/* 
   Declaration of variable:- var <identifier> type
   Example var number int32

   when a variable is declared with var it automatically initializes it to the 
   zero-value defined for its type. A type defines the set of values and the 
   set of operations that can take place on those values.

*/
package main

import "fmt"

func main(){

	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%T %T %T %T \n",i,f,b,s) // Print types of variable
	fmt.Printf("%v   %v    %v     %q \n",i,f,b,s) // Print initial values of variables
}

/*
Go construct and datatypes
----------------------------
Go file has a name or an identifier which is case sensitive like C.

The _(underscore) identifier is special. It is called blank identifier. It may be used in variable declarations.

It is like normal identifiers but its value is discarded, so it cannot be used anymore in the code.

It may happen that the variable, type, or function has no name and even enhance flexibility so it is called anonymous.

These are the 25 keywords for Go-code:
	break	default	func	interface	select
	case	defer	go	map	struct
	chan	else	goto	package	switch
	const	fallthrough	If	range	type
	continue	for	import	return	var

Additional
	append	bool	byte	cap	close	complex	complex64	complex128	uint16
	copy	false	float32	float64	imag	int	int8	int16	uint32
	int32	int64	iota	len	make	new	nil	panic	uint64
	print	println	real	recover	string	true	uint	uint8	Uintptr	
*/