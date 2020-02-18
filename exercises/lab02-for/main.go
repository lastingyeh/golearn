package main

import "fmt"

func main() {

	for i := 1; i <= 3; i++ {
		sum := 0.0
		for j := 1; j <= 5; j++ {
			score := 0.0
			fmt.Printf("The %dth class, %dth student score: \n", i, j)
			fmt.Scanln(&score)
			sum += score
		}
		fmt.Printf("The %dth class, avg: %f\n", i, sum/5.0)
	}
}
