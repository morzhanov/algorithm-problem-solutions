package array

import (
	"fmt"
)

// FindDuplicateElementInLimitedRangeArray simple solution
func FindDuplicateElementInLimitedRangeArray(arr []int) {
	sum := 0
	n := len(arr)
	for _, val := range arr {
		sum += val
	}
	expectedSum := n * (1 + n) / 2

	if sum != expectedSum {
		fmt.Printf("Duplicate: %v\n", n-(expectedSum-sum))
	} else {
		fmt.Printf("Duplicate not found\n")
	}
}
