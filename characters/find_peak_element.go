package characters

import (
	"fmt"
)

// FindPeakElement test
// A peak element is an element that is greater than its neighbors.
// Given an input array where num[i] ≠ num[i+1], find a peak element and return its index.
// The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.
func FindPeakElement(arr []int) {
	max := 0
	pos := 0
	// add first two elements to check first and last elements
	arr = append(arr, arr[0])
	arr = append(arr, arr[1])
	n := len(arr)

	for i := 1; i < n-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] && arr[i] > max {
			max = arr[i]
			pos = i % (n - 2)
		}
	}

	fmt.Printf("Peak element %v found on position %v\n", max, pos)
}
