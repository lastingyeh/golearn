package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var i int
	randNum := rand.Intn(100) + 1
	fmt.Println("correct num: ", randNum)
	var guessNum int

	for i < 10 {
		fmt.Printf("Guess %dth times: \n", (i + 1))
		fmt.Scanln(&guessNum)
		i++

		if randNum == guessNum {
			switch {
			case i == 1:
				fmt.Println("genius")
			case i < 3:
				fmt.Println("smart")
			case i < 9:
				fmt.Println("normal")
			default:
				fmt.Println("soso")
			}

			break
		}

		if i == 10 {
			fmt.Println("failure.")
		}
	}
}
