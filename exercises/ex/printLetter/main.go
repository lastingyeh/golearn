package main

import "fmt"

func main(){
	for i := 'a'; i <= 'z'; i ++ {
		fmt.Printf("%c ", i)
	}

	fmt.Println()

	for i := 'Z'; i >= 'A'; i-- {
		fmt.Printf("%c", i)
	}
}