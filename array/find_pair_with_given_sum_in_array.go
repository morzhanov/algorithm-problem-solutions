package array

import (
	"fmt"
	"sort"
)

// FindPairWithGivenSumInArray test
// given an unsorted array, find pair of values with givent sum of it
// prints indexes and values of those values
func FindPairWithGivenSumInArray(arr []int, sum int) {
	sort.Ints(arr)

	minIdx := 0
	maxIdx := len(arr) - 1

	for minIdx != maxIdx {
		min := arr[minIdx]
		max := arr[maxIdx]
		currentSum := min + max

		if currentSum == sum {
			fmt.Printf("Elements are found: %v + %v = %v\n", arr[minIdx], arr[maxIdx], sum)
			return
		}
		if max > sum {
			maxIdx--
		} else {
			minIdx++
		}
	}
	fmt.Print("Elements are not found")
}
