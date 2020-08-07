package array

import (
	"fmt"

	"github.com/morzhanov/algorithm-problem-solutions/utils"
)

// TrappingRainWatter test
// Given n non-negative integers representing an elevation map
// where the width of each bar is 1,
// compute how much water it is able to trap after raining.
// For example, given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.
func TrappingRainWatter(heights []int) {
	n := len(heights)
	res := 0
	resArr := make([]int, n)
	left := make([]int, n)
	right := make([]int, n)

	max := heights[0]
	left[0] = heights[0]
	for i := 1; i < n; i++ {
		if max < heights[i] {
			max = heights[i]
			left[i] = heights[i]
		} else {
			left[i] = max
		}
	}

	max = heights[n-1]
	right[n-1] = heights[n-1]
	for i := n - 2; i >= 0; i-- {
		if max < heights[i] {
			max = heights[i]
			right[i] = heights[i]
		} else {
			right[i] = max
		}
	}

	for i := 0; i < n; i++ {
		resArr[i] = utils.Min(left[i], right[i]) - heights[i]
		res += resArr[i]
	}

	fmt.Printf("Left: %v\n", left)
	fmt.Printf("Right: %v\n", right)
	fmt.Printf("Traps: %v, Water: %v\n", resArr, res)
}
