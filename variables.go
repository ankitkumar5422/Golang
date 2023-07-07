package main

import "fmt"

func main() {

	myvar1, myvar2 := 39, 45
	myvar3, myvar2 := 45, 100

	fmt.Print("The value of myvar1 and myvar2 is : \n", myvar1, myvar2)
	fmt.Print("The value of myvar3 and myvar2 is :%d %d\n", myvar3, myvar2)
}
