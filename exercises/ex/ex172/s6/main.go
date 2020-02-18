package main

import "fmt"

func main() {
	var arrStr [10]string
	makeArr(&arrStr)

	fmt.Println(arrStr)

	index := searchIndex(arrStr, "BB")
	fmt.Println(index)
}

func makeArr(arr *[10]string) {
	for i, j := 'A', 0; j < len(arr); i, j = i+1, j+1 {
		(*arr)[j] = fmt.Sprintf("%c%c", i, i)
	}
}

func searchIndex(arr [10]string, search string) (index int) {
	index = -1

	for i := 0; i < len(arr); i++ {
		if arr[i] == search {
			index = i
			return
		}
	}

	return
}
