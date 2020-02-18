package main

import (
	"code/exercises/ex/ex172/utils"
	"fmt"
)

const ln = 10

// 1. generate 10 random numbers
// reverse
// average
// max, maxIdx
// search val: 55

func main() {

	// 1. generate 10 random numbers
	arr := utils.IntSlice(ln)
	fmt.Println("slice: ", arr)

	// 2. reverse
	for i := 0; i < ln/2; i++ {
		arr[i], arr[ln-1-i] = arr[ln-1-i], arr[i]
	}

	fmt.Println("reverse: ", arr)

	// 3. avg
	// 4. max, maxIdx
	var sum int
	max, maxIdx := arr[0], 0

	for i := 1; i < ln; i++ {
		if max < arr[i] {
			max = arr[i]
			maxIdx = i
		}
		sum += arr[i]
	}

	fmt.Printf("max: %d, idx: %d, avg: %.2f\n", max, maxIdx, float64(sum)/ln)

	// 5 search
	// sorting by ASC
	// use binary search
	utils.Sort(arr, true)
	fmt.Println("after sorting: ", arr)

	binarySearch(arr, 0, ln, 55)
}

func binarySearch(arr []int, leftIdx, rightIdx, searchKey int) {
	if leftIdx > rightIdx {
		fmt.Printf("not found value: %d\n", searchKey)
		return
	}

	midIdx := (leftIdx + rightIdx) / 2

	if (arr)[midIdx] > searchKey {
		binarySearch(arr, leftIdx, midIdx-1, searchKey)
	} else if (arr)[midIdx] < searchKey {
		binarySearch(arr, midIdx+1, rightIdx, searchKey)
	} else {
		fmt.Printf("searchValue found, idx: %d\n", midIdx)
	}
}
