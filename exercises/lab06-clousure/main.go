package main

import (
	"fmt"
	"strings"
)

func main() {
	filename := makeSuffix(".jpg")("image.png")

	fmt.Println(filename)
}

func makeSuffix(suffix string) func(string) string {
	return func(filename string) string {
		if strings.HasSuffix(filename, suffix) {
			return filename
		}
		return filename + suffix
	}
}
