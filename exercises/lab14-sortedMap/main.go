package main

import (
	"fmt"
	"sort"
)

func main() {
	// sorted map
	// put key to slice
	// sorted slice
	// interate to map[slice[idx]]
	maps := map[int]string{3: "C", 1: "A", 8: "J", 6: "E"}

	var keys []int

	for key, _ := range maps {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	fmt.Println(keys)

	for _, idx := range keys {
		fmt.Printf("maps[%d] = %v\n", idx, maps[idx])
	}
}
