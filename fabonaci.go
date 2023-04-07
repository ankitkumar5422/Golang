package main

import "fmt"

func fibonacciSeries(num int) {

	a := 0
	b := 1
	c := b
	fmt.Printf("Fibonacci Series is: %d %d", a, b)
	for true {
		c = b
		b = a + b
		if b >= num {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)
	}
}
func main() {
	fibonacciSeries(5)

}
