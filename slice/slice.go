package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java", "Python"}
	fmt.Println(courseName)

	courseName = append(courseName, "C")
	fmt.Println(courseName)

	courseName = append(courseName, "C#", "PHP", "CSS", "HTML", "JavaScript")
	fmt.Println(courseName)

	courseWeb := courseName[5:8]
	fmt.Println(courseWeb)
	courseWeb = courseName[:4]
	fmt.Println(courseWeb)

}
