package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must number only", promt)
		panic(message)
	}

	return value
}

func getOperator() string {
	fmt.Print("operator is (+ - * /):")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func subtract(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}

func setwithpointer(str *string, opt string) {
	*str = opt
}

func main() {
	numbers := getInput("number of number u wanna calculate = ")
	var result float64
	var temp string
	var opt string

	for i := 0; i < int(numbers); i++ {

		value := getInput(" Number " + strconv.Itoa(i+1) + ": ")
		if i+1 < int(numbers) {
			opt = getOperator()
		}

		if i == 0 {
			result = value
			setwithpointer(&temp, opt)
			continue
		}

		switch temp {
		case "+":
			result = add(result, value)
		case "-":
			result = subtract(result, value)
		case "*":
			result = multiply(result, value)
		case "/":
			result = divide(result, value)
		default:
			panic("wrong operation")
		}
		setwithpointer(&temp, opt)

	}

	fmt.Printf("result = %v", result)
}
