package main

import (
	"code/exercises/ex/ex172/utils"
	"fmt"
	"math"
)

func main() {
	// referee 8 -> scores 8
	// exclude by largest & smallest -> avg(6)
	// find largest & smallest
	// find better referee scores -> close to abs(avg(6))
	// find worse referee scores -> far to abs(avg(6))

	scores := utils.IntSlice(8)
	fmt.Println("scores: ", scores)
	// include max, min and caculate avg
	max, min, sum := scores[0], scores[0], 0

	for i := 1; i < len(scores); i++ {
		if max < scores[i] {
			max = scores[i]
		} else if min > scores[i] {
			min = scores[i]
		}
		sum += scores[i]
	}

	fmt.Printf("max: %d, min: %d\n", max, min)

	avg := float64(sum-max-min) / float64(len(scores)-2)

	fmt.Printf("avg: %.2f\n", avg)

	bIdx, wIdx := 0, 0
	bDiff, wDiff := math.Abs(float64(scores[0])-avg), math.Abs(float64(scores[0])-avg)

	fmt.Printf("absDiff(0): %.2f\n", bDiff)

	for i := 1; i < len(scores); i++ {
		absDiff := math.Abs(float64(scores[i]) - avg)

		fmt.Printf("absDiff(%d): %.2f\n", i, absDiff)

		if wDiff < absDiff {
			wDiff = absDiff
			wIdx = i
		} else if bDiff > absDiff {
			bDiff = absDiff
			bIdx = i
		}
	}

	fmt.Printf("better: %d, worse: %d\n", bIdx, wIdx)
}
