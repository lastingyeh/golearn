package main

import "fmt"

func main() {

	var i int = 1
	var count int
	var subSum int

	for i <= 100 {
		if isPrimeNum(i) {

			count++

			if count%6 != 0 {
				subSum += i
				fmt.Printf("%d\t", i)

			} else {
				fmt.Printf("sum = %d", subSum)
				subSum = 0
				fmt.Println()
			}
		}

		if i == 100 {
			fmt.Printf("sum = %d\n", subSum)
		}
		
		i++
	}
}

func isPrimeNum(n int) bool {
	for i := 2; i <= n; i++ {
		if n%i == 0 && n != i {
			return false
		}
	}
	return true
}
