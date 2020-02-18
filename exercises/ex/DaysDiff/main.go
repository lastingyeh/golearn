package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.Local).Unix()
	now := time.Now().Unix()

	dayDiff := (now - start) / 60 / 60 / 24

	switch dayDiff % 5 {
	case 1, 2, 3:
		fmt.Println("fishing")
	default:
		fmt.Println("rest")
	}

	//fmt.Println(now.Sub(start).)
}
