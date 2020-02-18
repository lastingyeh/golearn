package main

import "fmt"

func main() {
	fibs := fib(10)
	fmt.Println(fibs)
}

func fib(n int) []uint64 {
	var fibs = make([]uint64, n)

	fibs[0] = 1
	fibs[1] = 1

	for i := 2; i < n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}

	return fibs
}
