package main

import (
	"fmt"
)

// LongestValidParentheses test
// given a string containing just the characters '(' and ')',
// find the length of the longest valid (well-formed) parentheses substring.
func LongestValidParentheses(arr []int) {
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
