package main

import "fmt"

func main() {
	// 1111       	4444
	// 2222    => 	3333
	// 3333       	2222
	// 4444       	1111

	// i = 0 <-> i = 3
	// i = 1 <-> i = 2
	var d2arr [4][4]int

	make2darr(&d2arr)

	print(d2arr)

	// reverse row
	for i := 0; i < len(d2arr)/2; i++ {
		for j := 0; j < len(d2arr[i]); j++ {
			d2arr[i][j], d2arr[len(d2arr)-1-i][j] = d2arr[len(d2arr)-1-i][j], d2arr[i][j]
		}
	}

	print(d2arr)

}

func make2darr(d2arr *[4][4]int) {
	for i := 0; i < len(d2arr); i++ {
		for j := 0; j < len(d2arr[i]); j++ {
			(*d2arr)[i][j] = i
		}
	}
}

func print(arr [4][4]int) {
	for _, d1 := range arr {
		for _, d2 := range d1 {
			fmt.Print(d2)
		}
		fmt.Println()
	}

	fmt.Println()
}
