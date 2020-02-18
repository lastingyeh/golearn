package main

import "fmt"

func main() {
	// func params: arr
	// find max, maxIdx, min, minIdx
	arr := [...]int{95, 10, 38, 22, 56}
	maxIdx, minIdx := arrMaxMin(arr)

	fmt.Printf("maxIdx: %d, minIdx: %d\n", maxIdx, minIdx)
}

func arrMaxMin(arr [5]int) (maxIdx, minIdx int) {
	max, min := arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			maxIdx = i
		} else if arr[i] < min {
			min = arr[i]
			minIdx = i
		}
	}
	return
}
