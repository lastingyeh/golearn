package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var arr [5]int

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	fmt.Println("originate = ", arr)
	reverse((&arr))
	fmt.Println("reverse = ", arr)
}

func reverse(arr *[5]int) {

	// 0 1 2 3 4
	// 5 0
	ln := len(arr) / 2
	// var tmp int

	for i := 0; i < ln; i++ {
		// tmp = (*arr)[len(arr)-1-i]
		// (*arr)[len(arr)-1-i] = (*arr)[i]
		// (*arr)[i] = tmp
		(*arr)[i], (*arr)[len(arr)-1-i] = (*arr)[len(arr)-1-i], (*arr)[i]
	}

}
