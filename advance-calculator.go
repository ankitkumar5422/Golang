package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	value1 := getInput("value 1")
	value2 := getInput("value2")

	var result float64

	switch operation := getOperation(); operation {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subtractValues(value1, value2)
	case "*":
		result = multipleValues(value1, value2)
	case "/":
		result = divideValues(value1, value2)
	default:
		panic("Invalid operation")
	}
	fmt.Println("The result is %v:\n", result)

}
func getInput(prompt string) float64 {
	fmt.Printf("%v:", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		message := fmt.Sprintf("%v must be a number", prompt)
		panic(message)
	}
	return value
}

func getOperation() string {
	fmt.Print("Select an operation ( + - * /):")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}
func addValues(value1, value2 float64) float64 {
	return value1 + value2
}
func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}
func multipleValues(value1, value2 float64) float64 {
	return value1 * value2
}
func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
