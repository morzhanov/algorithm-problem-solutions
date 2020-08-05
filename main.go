package main

import (
	"fmt"
	"math/rand"
)

func createArray() []int {
	values := make([]int, 20)
	for i := range values {
		values[i] = rand.Intn(100)
	}
	return values
}

func main() {
	arr := [10]int{1, 2, 4, 5, 2, 3, 5, 1, 2, 4}
	FindPairWithGivenSumInArray(arr[:], 8)
	fmt.Printf("\n")

	arr = [10]int{1, -4, 4, 2, 2, -2, 4, 5, -8, -4}
	FindAllSubarraysWithZeroSum(arr[:])
	fmt.Printf("\n")

	arr = [10]int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9}
	FindDuplicateElementInLimitedRangeArray(arr[:])
	fmt.Printf("\n")

	str := "(){{{{}}}}"
	ValidateParentheses(str)
	str = "{{{({}){([()])}}}}"
	ValidateParentheses(str)
	str = "{{{{}]}}}"
	ValidateParentheses(str)
	fmt.Printf("\n")

	str = "(((())))(((((())))))"
	LongestValidParentheses(str)
	str = "(((()))))(())"
	LongestValidParentheses(str)
	str = "(((())))(((())))"
	LongestValidParentheses(str)
	fmt.Printf("\n")

	str = "Red rum sir is murder"
	ValidPalindrome(str)
	str = "Red rum sir is not murder"
	ValidPalindrome(str)
	str = ""
	ValidPalindrome(str)
	fmt.Printf("\n")

	str = "qwertyghqydidapjdwoasodwja"
	LongestSubstringWithoutRepeatingCharacters(str)
	fmt.Printf("\n")
}
