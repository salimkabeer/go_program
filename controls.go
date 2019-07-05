/*

*/
if(boolean_expression) {  
   /* statement(s) got executed only if the expression results in true */  
} else {  
   /* statement(s) got executed only if the expression results in false */  
}
---------------------------------------------------------------------------
The switch statement in Go is more flexible. In the above syntax, var1 
is a variable which can be of any type, and val1, val2, ... are possible values of var1.

switch  var1 {  
case val1:  
.....  
case val2  
.....  
default:  
.....  
}  

In switch statement, more than one values can be tested in a case, the values
are presented in a comma separated list. like: case val1, val2, val3:

If any case is matched, the corresponding case statement is executed. Here, 
the break keyword is implicit. So automatic fall-through is not the default 
behavior in Go switch statement.

For fall-through in Go switch statement, use the keyword "fallthrough" at 
the end of the branch.

-------------------------------------------------------------------------------

The Go for statement is used for repeating a set of statements number of times. 
It is the only loop in go language. There are two variants of for loop in Go: 
		Counter-controlled iteration and Condition-controlled iteration.

When the execution of the loop is over, the objects created inside the loop gets destroyed.

for a := 0; a < 5; a++ {  
      fmt.Print(a,"\n")
}

for i:=0; ; i++ // infinite loop
for {}  // infinite loop

for sum < 100 {  
      sum += sum  
      fmt.Println(sum)  
   }

-------------------------------------------------------------------------------
// goto statement program

package main  
import (  
   "fmt"  
)  
func main() {  
   var input int  
Loop:  
   fmt.Print("You are not eligible to vote ")  
   fmt.Print("Enter your age ")  
   fmt.Scanln(&input)  
   if (input <= 17) {  
      goto Loop  
   } else {  
      fmt.Print("You can vote ")  
   }  
}  

-------------------------------------------------------------------------------
A break statement is used to break out of the innermost structure in which it 
occurs. It can be used in for-loop (counter, condition,etc.), and also in a 
switch. Execution is continued after the ending } of that structure.

-------------------------------------------------------------------------------
The continue is used to skip the remaining part of the loop, and then continues 
with the next iteration of the loop after checking the condition.

-------------------------------------------------------------------------------
// single line comments
/* multi-line
   comments
*/

-------------------------------------------------------------------------------
// Constants
A constant const contains data which is not changed. This data can only be of 
type boolean, number (integer, float or complex) or string.

Syntax:- const identifier [type] = value 
Example:- const PI = 3.14
	  const str string = "salim"

-------------------------------------------------------------------------------
Type casting means conversion of a variable from one data type to another. 
The value may be lost when large type is converted to a smaller type.

-------------------------------------------------------------------------------
			THE END
