package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var pwdPath string = "c:/Users/61101/Documents/R/Go-practice/file/"

func main() {
	file, err := os.Open(pwdPath + "indexinfo.csv")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		item := strings.Split(line, ",")
		fmt.Println(item[3])
		// fmt.Printf("Line %d : %s \n", count, line)

		count++
	}

}
