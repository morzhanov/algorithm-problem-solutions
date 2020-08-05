package characters

import (
	"fmt"
)

// FindAllSubarraysWithZeroSum test
// given an unsorted array, find all subarrays with 0 sum
func FindAllSubarraysWithZeroSum(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			if sum == 0 {
				fmt.Printf("Subarray found: %v...%v\n", i, j)
			}
		}
	}
}
