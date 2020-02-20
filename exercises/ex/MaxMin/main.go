package main

import "fmt"

func main() {
	arr := []int{30, 20, 7, 6, 5, 4, 99}
	max, min := getMaxAndMin(arr)

	fmt.Println(max, min)
}

func getMaxAndMin(arr []int) (max, min int) {

	// s: small, b: big
	// separate by 2 parts ([{s, b}, {s, b}, {s, b}])
	for i := 0; i < len(arr)-1; i = i + 2 {
		if arr[i] > arr[i+1] {
			tmp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = tmp
		}
	}

	min = arr[0] 
	max = arr[1] 

	fmt.Println(arr)
	// 's' part of array compare to min
	for i := 2; i < len(arr); i = i + 2 {
		if min > arr[i] {
			min = arr[i]
		}
	}
	// 'b' part of array compare to max
	for i := 3; i < len(arr); i = i + 2 {
		if max < arr[i] {
			max = arr[i]
		}
	}

	// check last as arr is odd numbers array
	if len(arr)%2 == 1 {
		if max < arr[len(arr)-1] {
			max = arr[len(arr)-1]
		}

		if min > arr[len(arr)-1] {
			min = arr[len(arr)-1]
		}
	}

	return
}
