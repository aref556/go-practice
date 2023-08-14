package main

import "fmt"

var numberInt, numberInt2 int = 1000, 2000
var msg string = "Hello world"

func main() {

	numberfloat := 23.5

	fmt.Printf("%T \n", numberfloat)
	fmt.Println(numberInt)
	fmt.Println(numberInt2)
	fmt.Println(numberInt2 + numberInt)
	fmt.Println(numberInt + int(numberfloat))
	fmt.Println(numberInt + msg)

}
