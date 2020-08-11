package array

import (
	"fmt"
)

// MinimumSizeSubarraySum test
// Given an array of n positive integers and a positive integer s,
// find the minimal length of a subarray of which the sum ≥ s.
// If there isn't one, return 0 instead.
// For example, given the array [2,3,1,2,4,3] and s = 7,
// the subarray [4,3] has the minimal length of 2 under the problem constraint.
func MinimumSizeSubarraySum(arr []int, s int) {
	n := len(arr)

	res := n
	sum := 0
	start := 0
	i := 0
	for i != n {
		fmt.Printf("%v %v\n", start, i)
		if sum >= s {
			sum -= arr[start]
			start++
			if res > i-start {
				res = i - start
			}
		} else {
			if i == n {
				break
			}
			sum += arr[i]
			i++
		}
	}
	fmt.Printf("Minimun subarray with sum %v is %v length", s, res)
}
