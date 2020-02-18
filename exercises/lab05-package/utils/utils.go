package utils

import "fmt"

var Age int
var Name string

func init() {
	Age = 18
	Name = "NoName"

	fmt.Println("utils:init")
}

