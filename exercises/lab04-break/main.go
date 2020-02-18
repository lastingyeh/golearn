package main

import "fmt"

func main() {
	PathCount(60000)
}

func scanNumberCount() {
	var pCount int
	var nCount int
	var num int

	for {
		fmt.Println("Enter number: ")
		fmt.Scanln(&num)

		if num == 0 {
			break
		}

		if num > 0 {
			pCount++
			continue
		}

		nCount++
	}

	fmt.Printf("postive count: %d, negative count: %d", pCount, nCount)
}

func PathCount(money float64) {
	var pCount int

	for {
		fmt.Println(money)
		if money > 50000 {
			money = money * 0.95
		} else {
			money -= 1000
		}
		pCount++

		if money < 1000 {
			break
		}
	}
	fmt.Println("path count: ", pCount)
}
