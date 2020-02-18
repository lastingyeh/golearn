package main

import "fmt"

func main() {
	var d2arr [3][4]int
	var num int
	// i = 0
	// j = 0, j = len(j)
	// 0000 i=[0:0]
	// 0120 i=[1,0]
	// 0000 i=[2,0]

	for i := 0; i < len(d2arr); i++ {
		for j := 0; j < len(d2arr[i]); j++ {
			if !(i == 0 || i == len(d2arr)-1 || j == 0 || j == len(d2arr[i])-1) {
				fmt.Println("number: ")
				fmt.Scanln(&num)
				d2arr[i][j] = num
			}
		}
	}

	print(d2arr)
}

func print(arr [3][4]int) {
	for _, d1 := range arr {
		for _, d2 := range d1 {
			fmt.Print(d2)
		}
		fmt.Println()
	}
}
