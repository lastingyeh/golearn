package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func IntSlice(size int) []int {

	var slice = make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100) + 1
	}
	return slice
}

func Sort(arr []int, isASC bool) {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if isASC {
				if (arr)[j] > (arr)[j+1] {
					(arr)[j], (arr)[j+1] = (arr)[j+1], (arr)[j]
				}
			} else {
				if (arr)[j] < (arr)[j+1] {
					(arr)[j], (arr)[j+1] = (arr)[j+1], (arr)[j]
				}
			}
		}
	}
}
