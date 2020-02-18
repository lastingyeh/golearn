package main

import (
	"code/exercises/ex/ex172/utils"
	"fmt"
)

// given sorting array (ASC)
// insert value to array
// result: still sorting array (ASC)

func main() {
	const ln = 10
	arr := utils.IntSlice(ln)
	fmt.Println("1. ", arr)

	// sorting
	utils.Sort(arr, true)
	fmt.Println("2. ", arr)

	// target insert number: 49
	nArr := insertToArr(arr, 0, ln, 49)

	fmt.Println("3. ", nArr)
}

func insertToArr(arr []int, leftIdx, rightIdx, insertValue int) []int {
	var nArr []int
	
	// got return condition
	if leftIdx > rightIdx {
		nArr = append(nArr, append(arr[:leftIdx], insertValue)...)
		nArr = append(nArr, arr[leftIdx:]...)

		// fmt.Println("3. ", nArr)

		return nArr
	}

	mid := (leftIdx + rightIdx) / 2 // 5

	if arr[mid] > insertValue {
		nArr = insertToArr(arr, leftIdx, mid-1, insertValue)
	} else {
		nArr = insertToArr(arr, mid+1, rightIdx, insertValue)
	}

	return nArr
}
