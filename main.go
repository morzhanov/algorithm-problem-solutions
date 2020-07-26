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

	arr2 := [20]int{1, -4, 4, 2, 2, -2, 4, 5, -8, -4, 2, 3, 5, -6, 1, 2, -9, -2, 4, 5}
	FindAllSubarraysWithZeroSum(arr2[:])
	fmt.Printf("\n")

	arr3 := [8]int{1, 2, 3, 4, 4, 5, 6, 7}
	FindDuplicateElementInLimitedRangeArray(arr3[:])
	fmt.Printf("\n")

	str := "(){{{{}}}}"
	str2 := "{{{({}){([()])}}}}"
	str3 := "{{{{}]}}}"
	ValidateParentheses(str)
	ValidateParentheses(str2)
	ValidateParentheses(str3)
	fmt.Printf("\n")

	str4 := "(((())))(((())))"
	str5 := "(((()))))(())"
	str6 := "(((())))(((((())))))"
	LongestValidParentheses(str4)
	LongestValidParentheses(str5)
	LongestValidParentheses(str6)
	fmt.Printf("\n")
}
