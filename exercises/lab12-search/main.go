package main

import "fmt"

// general search
func search(code string) int {
	codings := [4]string{"Java", "Golang", "C#", "Python"}
	var idx = -1

	for i := 0; i < len(codings); i++ {
		if codings[i] == code {
			idx = i
		}
	}
	return idx
}

// binary serach
// 1. sorted (ASC)
// 2. find middle ((rightIdx + leftIdx) / 2)
// 3.1 arr[middle] > findVal -> arr[leftIdx : middle-1]
// 3.2 arr[middle] < findVal -> arr[middle+1 : rightIdx]
// 3.3 arr[middle] = findVal -> arr[middle]
// 4. if leftIdx > rightIdx return -> not found
func binarySearch(arr []int, leftIdx int, rightIdx int, searchVal int) {
	// return condition
	if leftIdx > rightIdx {
		fmt.Println("not found")
		return
	}

	// get midIdx
	midIdx := (rightIdx + leftIdx) / 2
	// arr[leftIdx:midIdx-1]
	if arr[midIdx] > searchVal {
		binarySearch(arr, leftIdx, midIdx-1, searchVal)
		// arr[midIdx+1:rightIdx]
	} else if arr[midIdx] < searchVal {
		binarySearch(arr, midIdx+1, rightIdx, searchVal)
		// find
	} else {
		fmt.Printf("found value: %d, idx: %d\n", searchVal, midIdx)
	}
}

func main() {
	var searchStr = ""

	fmt.Println("Which codings you selected.")
	fmt.Scanln(&searchStr)

	if idx := search(searchStr); idx != -1 {
		fmt.Printf("coding: %v, idx: %v\n", searchStr, idx)
	}

	var searchVal int
	fmt.Scanln(&searchVal)
	arr := []int{1, 8, 10, 89, 100, 1234}

	binarySearch(arr, 0, len(arr)-1, searchVal)
}
