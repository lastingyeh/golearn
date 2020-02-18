package main

import "fmt"

func main() {
	// arr [1, 3, 5, 7, 9]
	// arr [9, 7, 5, 3, 1]

	var arr [5]int
	i, j := 0, 0

	for {
		if i%2 == 1 {
			arr[j] = i
			j++
		}
		i++

		if len(arr) == j {
			break
		}
	}

	fmt.Println(arr)

	// reverse
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Println(arr)
}
