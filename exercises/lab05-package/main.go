package main

import (
	"code/exercises/lab05-package/utils"
	"fmt"
)

func init(){
	fmt.Println("main:init")
}

func main() {
	fmt.Printf("main:main: Init Age: %d, Name: %s", utils.Age, utils.Name)
}
