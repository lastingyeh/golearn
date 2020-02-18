package main

import "fmt"

func bubbleSorting(arr *[5]int) {
	// r1
	// for j := 0; j < len((*arr))-1; j++ {
	// 	if (*arr)[j] > (*arr)[j+1] {
	// 		(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
	// 	}
	// }

	// r2
	// for j := 0; j < len((*arr))-2; j++ {
	// 	if (*arr)[j] > (*arr)[j+1] {
	// 		(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
	// 	}
	// }

	// r3
	// .... j < len((*arr))-3;...

	// r4
	// ... j < lne((*arr))-4; ...

	// summary
	for i := 0; i < len((*arr))-1; i++ {
		for j := 0; j < len((*arr))-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}

}

func main() {
	arr := [...]int{24, 69, 80, 57, 13}

	fmt.Println("before sorting: ", arr)

	bubbleSorting(&arr)

	fmt.Println("after sorting: ", arr)
}
