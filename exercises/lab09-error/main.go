package main

import (
	"errors"
	"fmt"
)

func main() {
	divideByZero()
	fmt.Println("done")

	readConf()
}

func divideByZero() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var n int = 10
	var m int = 0
	res := n / m
	fmt.Println("res = ", res)
}

func customError(filename string) error {
	if filename == "config.ini" {
		return nil
	}
	return errors.New("read filename error")
}

func readConf(){
	err := customError("config.in")
	if err != nil {
		panic(err)
	}
	fmt.Println("read successful.")
}
