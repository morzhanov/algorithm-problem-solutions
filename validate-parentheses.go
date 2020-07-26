package main

import "fmt"

// ValidateParentheses test
// given a string containing just the characters '(){}[]'
// return true if string has valid open-closed sequence of characters
func ValidateParentheses(str string) {
	arr := []rune(str)
	n := len(arr)
	chars := [6]rune{'(', ')', '{', '}', '[', ']'}
	stack := make([]int, n)

	for _, char := range arr {
		idx := Pos(chars[:], char)

		if idx == -1 {
			fmt.Printf("String is not valid\n")
			return
		}

		if idx%2 == 0 {
			stack = append(stack, idx+1)
		} else {
			stackN := len(stack)
			curr := stack[stackN-1]
			stack = stack[:stackN-1]
			if curr != idx {
				fmt.Printf("String is not valid\n")
				return
			}
		}
	}
	fmt.Printf("String is valid\n")
}
