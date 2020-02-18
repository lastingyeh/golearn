package main

import "fmt"

func main() {
	// arr [8]int
	// count: i > avg
	// count: i < avg
	arr := [...]int{1, 3, 9, 88, 34, 23, 77, 56}

	ngc, ltc := countAvgNums(arr)

	fmt.Printf("ngc: %d, ltc: %d\n", ngc, ltc)
}

func countAvgNums(arr [8]int) (ngc, ltc int) {
	var avg float64
	var sum int

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	avg = float64(sum) / float64(len(arr))

	fmt.Println("avg: ", avg)

	for i := 0; i < len(arr); i++ {
		if float64(arr[i]) > avg {
			ngc++
		} else if float64(arr[i]) < avg {
			ltc++
		}
	}
	return
}
