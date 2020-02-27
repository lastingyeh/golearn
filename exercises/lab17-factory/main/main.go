package main

import (
	"code/exercises/lab17-factory/model"
	"fmt"
)

func main() {
	var stu = model.NewStudent("Tom", 78.9)
	fmt.Printf("student = %v\n", stu)

	// set name & score
	stu.SetName("Jack")
	stu.SetScore(99)

	fmt.Printf("get name = %s\n", stu.GetName())
	fmt.Printf("get score = %.2f\n", stu.GetScore())

	fmt.Printf("student = %v\n", stu)
}
