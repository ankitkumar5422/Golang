// Go program to illustrate the
// global variables
package main

import "fmt"

// global variable declaration
var myvariable1 int = 100

func main() { // from here local level scope starts

	// local variables inside the main function
	var myvariable2 int = 200

	// Display the value of global variable
	fmt.Printf("The value of Global myvariable1 is : %d\n",
		myvariable1)

	// Display the value of local variable
	fmt.Printf("The value of Local myvariable2 is : %d\n",
		myvariable2)

	// calling the function
	display()

} // here local level scope ends

// taking a function
func display() { // local level starts

	// Display the value of global variable
	fmt.Printf("The value of Global myvariable1 is : %d\n",
		myvariable1)

} // local scope ends here
