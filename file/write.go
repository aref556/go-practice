package main

import (
	"fmt"
	"os"
)

var pwdPath string = "c:/Users/61101/Documents/R/Go-practice/file/"

// func check(err error) {
// 	panic(err)
// }

func main() {
	data1 := []byte("Hi \n Helloworld")
	err := os.WriteFile(pwdPath+"data.txt", data1, 0644)
	if err != nil {
		fmt.Println("WriteFile")
		panic(err)
	}

	f, err := os.Create(pwdPath + "employeeName")
	if err != nil {
		fmt.Println("Create")
		panic(err)
	}

	defer f.Close()

	data2 := []byte("Sira\n Manee")
	os.WriteFile(pwdPath+"employeeName.txt", data2, 0644)
}
