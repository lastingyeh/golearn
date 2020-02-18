package main

import "fmt"

func main() {
	// new: allocate value type
	var n = new(int)
	inc(n)
	fmt.Printf("Type: %T, value: %d, addr: %d, pValue: %d\n", n, n, &n, *n)

	// make: allocate ref type: channel, slice, map
}

func inc(n *int) {
	*n++
}
