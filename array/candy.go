package array

import (
	"fmt"
)

// Candy test
// here are N children standing in a line. Each child is assigned a rating value.
// You are giving candies to these children subjected to the following requirements:
//	1. Each child must have at least one candy.
//	2. Children with a higher rating get more candies than their neighbors.
// What is the minimum candies you must give?
func Candy(ratings []int) {
	n := len(ratings)
	res := 0
	resArr := make([]int, n)

	for i := 1; i < n; i++ {
		curr := 1
		if ratings[i] > ratings[i-1] {
			curr = resArr[i-1] + 1
		}
		resArr[i] = curr
		res += curr
	}

	for i := n - 2; i >= 0; i-- {
		curr := 1
		if ratings[i] > ratings[i+1] {
			curr = resArr[i+1] + 1
		}
		resArr[i] = curr
		res += curr
	}

	fmt.Printf("Ratings: %v. Candies: %v. Minimum amount of candies: %v\n", ratings, resArr, res)
}
