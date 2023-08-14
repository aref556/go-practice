package main

import "fmt"

var score int

func main() {

	fmt.Println("Grade Calucaltir")
	fmt.Scanf("%d", &score)
	fmt.Println(" score = ", score)

	if score >= 80 {
		fmt.Println("Your grade is: A")
	} else if score >= 70 {
		fmt.Println("Your grade is: B")
	} else if score >= 60 {
		fmt.Println("Your grade is: C")
	} else if score >= 50 {
		fmt.Println("Your grade is: D")
	} else {
		fmt.Println("Your grade is: F")
	}
}
