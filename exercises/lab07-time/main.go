package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// current time
	now := time.Now()
	fmt.Printf("time type = %T, value = %v", now, now)

	// get year / month / day / hr / min /sec
	fmt.Printf("year = %v\n", now.Year())
	fmt.Printf("month = %v (%v)\n", now.Month(), int(now.Month()))
	fmt.Printf("day = %v\n", now.Day())
	fmt.Printf("hours = %v\n", now.Hour())
	fmt.Printf("minutes = %v\n", now.Minute())
	fmt.Printf("seconds = %v\n", now.Second())

	// format time 1
	fmt.Printf("%d-%d-%d %d:%d:%d\n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	// format time 2 (fixed value for params)
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println()

	fmt.Println(now.Format("2006-01"))
	fmt.Println()

	// constant
	var n int

	for n <= 5 {
		n++
		fmt.Println(n)
		time.Sleep(time.Millisecond * 100)
	}

	// Unix & UnixNano on rand seed usage from 1970-01-01
	fmt.Println("Unix: ", time.Now().Unix())
	fmt.Println("UnixNano: ", time.Now().UnixNano())

	// exec duration
	start := time.Now().Unix()
	procTime()
	end := time.Now().Unix()

	fmt.Printf("work duaration: %v secs\n", end-start)

}

// exercise
func procTime() {
	str := ""

	for i := 0; i < 100000; i++ {
		str += strconv.Itoa(i)
	}
}
