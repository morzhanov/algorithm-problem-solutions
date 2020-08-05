package str

import (
	"fmt"
)

// LongestValidParentheses test
// given a string containing just the characters '(' and ')',
// find the length of the longest valid (well-formed) parentheses substring.
func LongestValidParentheses(str string) {
	arr := []rune(str)
	stack := make([][]int, 0)
	res := 0

	for i, char := range arr {

		if char == '(' {
			stack = append(stack, []int{0, i})
		} else {
			if len(stack) != 0 && stack[len(stack)-1][0] == 0 {
				stack = stack[:len(stack)-1]
				curr := 0
				if len(stack) == 0 {
					curr = i + 1
				} else {
					curr = stack[len(stack)-1][1]
				}
				if curr > res {
					res = curr
				}
			} else {
				stack = append(stack, []int{1, i})
			}
		}
	}

	fmt.Printf("Longest valid distance: %v\n", res)
}
