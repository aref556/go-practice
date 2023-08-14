package main

import "fmt"

var productName [4]string

// var price [4]float32

func main() {
	productName[0] = "MacBook"
	productName[0] = "iPad"
	productName[0] = "iPhone"
	productName[0] = "AirPods"

	price := [4]float32{40000, 30000, 20000, 2000}

	fmt.Println(price)
}
