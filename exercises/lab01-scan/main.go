package main

import (
	"fmt"
)

type user struct {
	Name   string
	Age    byte
	Salary float32
	Pass   bool
}

func Scanln(u *user) {
	fmt.Println("Name:")
	fmt.Scanln(&u.Name)

	fmt.Println("Age:")
	fmt.Scanln(&u.Age)

	fmt.Println("Salary:")
	fmt.Scanln(&u.Salary)

	fmt.Println("Pass:")
	fmt.Scanln(&u.Pass)
}

func Scanf(u *user) {
	fmt.Println("Enter name, age, salary, pass separated by space")
	fmt.Scanf("%s %d %f %t", &u.Name, &u.Age, &u.Salary, &u.Pass)
}

func main() {
	user1 := &user{}
	Scanln(user1)

	fmt.Println("user1: ", user1)

	user2 := &user{}
	Scanf(user2)

	fmt.Println("user2: ", user2)
}
