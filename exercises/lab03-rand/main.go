package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now().Unix()
	// unix time
	fmt.Println(t)
	// set rand seed by time
	rand.Seed(t)

	count := 0
	// check rand number == 99 break
	for {
		n := rand.Intn(100) + 1
		count++
		if n == 99 {
			break
		}
	}

	fmt.Println(count)
}
