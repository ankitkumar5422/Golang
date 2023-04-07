package main

import "fmt"

func fibonacciSeries(n int) {

	a := 0
	b := 1

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)

		c := a + b
		a = b
		b = c

	}

}
func main() {
	fibonacciSeries(5)

}
