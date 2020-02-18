package main

import (
	"fmt"
)

func main() {
	for {
		var year int
		var month int
		var day int
		var errFlag bool
		// scan year, month, day
		fmt.Println("Year: ")
		fmt.Scanln(&year)

		fmt.Println("month: ")
		fmt.Scanln(&month)

		fmt.Println("day: ")
		fmt.Scanln(&day)

		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			if day > 31 {
				errFlag = true
			}
		case 2:
			if isLeapYear(year) {
				if day > 29 {
					errFlag = true
				}
			} else if day > 28 {
				errFlag = true
			}
		case 4, 6, 9, 11:
			if day > 30 {
				errFlag = true
			}
		}

		if errFlag {
			fmt.Println("date value error")
			continue
		} else {
			fmt.Printf("%d-%d-%d\n", year, month, day)
		}
	}
}

func isLeapYear(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}
