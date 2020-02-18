package main

import (
	"code/exercises/ex/ex172/utils"
	"fmt"
)

func main() {
	// generate 10 numbers of slice
	// bubble sorting
	// binarySearch

	nums := utils.IntSlice(10)
	fmt.Println("1. ", nums)

	utils.Sort(nums, true)
	fmt.Println("2. ", nums)

	binarySearch(nums, 0, len(nums)-1, 90)
}

func binarySearch(arr []int, lIdx, rIdx, searchValue int) {
	if lIdx > rIdx {
		fmt.Println("not found.")
		return
	}

	mid := (lIdx + rIdx) / 2

	if arr[mid] > searchValue {
		binarySearch(arr, lIdx, mid-1, searchValue)
	} else if arr[mid] < searchValue {
		binarySearch(arr, mid+1, rIdx, searchValue)
	} else {
		fmt.Println("search idx: ", mid)
	}
}
