package dynamic

import "fmt"

// HouseRobber test
// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed,
// the only constraint stopping you from robbing each of them
// is that adjacent houses have security system connected
// and it will automatically contact the police
// if two adjacent houses were broken into on the same night.
//
// Given a list of non-negative integers representing the amount of money of each house,
// determine the maximum amount of money you can rob tonight without alerting the police.
func HouseRobber(arr []int) {
	houses := make([]int, 0)

	if len(arr) == 0 {
		fmt.Printf("Max amount of money to rob is: %v\n", 0)
	}
	if len(arr) == 1 {
		fmt.Printf("Max amount of money to rob is: %v\n", arr[0])
	}

	if arr[0] > arr[1] {
		houses = append(houses, arr[0])
	} else {
		houses = append(houses, arr[1])
	}

	for i := 2; i < len(arr); i++ {
		curr := arr[i-1]
		planned := arr[i] + houses[i-2]
		if planned > curr {
			houses = append(houses, planned)
		} else {
			houses = append(houses, curr)
		}
	}

	fmt.Printf("Max amount of money to rob is: %v\n", houses[len(houses)-1])
}
