package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

// implement sort methods
type Heros []Hero

// Len is the number of elements in the collection.
// Len() int
func (h Heros) Len() int {
	return len(h)
}

// Less reports whether the element with
// index i should sort before the element with index j.
// Less(i, j int) bool
func (h Heros) Less(i, j int) bool {
	return h[i].Age < h[j].Age
}

// Swap swaps the elements with indexes i and j.
// Swap(i, j int)
func (h Heros) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	arr := []int{1, -5, 9, 8, 30}
	sort.Ints(arr)

	fmt.Println(arr)

	// sort by hero's age
	// generate heros slice
	var heros Heros
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("hero~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}

	fmt.Println("before sorting...")

	for _, v := range heros {
		fmt.Println(v)
	}

	fmt.Println("after sorting...")

	sort.Sort(heros)

	for _, v := range heros {
		fmt.Println(v)
	}
}
