package main

import (
	"fmt"
	"strconv"
)

func fooer(input int) string {
	isfoo :=(input %3)==0

	if isfoo {
		return "Foo"
	}
	return strconv.Itoa(input)
}

func main (){
	input := 3
	fmt.Println(fooer(input))

}