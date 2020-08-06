package main

import (
	"fmt"
	"math/rand"

	"github.com/morzhanov/algorithm-problem-solutions/characters"
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
	characters.FindPairWithGivenSumInArray(arr[:], 8)
	fmt.Printf("\n")

	arr = [10]int{1, -4, 4, 2, 2, -2, 4, 5, -8, -4}
	characters.FindAllSubarraysWithZeroSum(arr[:])
	fmt.Printf("\n")

	arr = [10]int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9}
	characters.FindDuplicateElementInLimitedRangeArray(arr[:])
	fmt.Printf("\n")

	str := "(){{{{}}}}"
	characters.ValidateParentheses(str)
	str = "{{{({}){([()])}}}}"
	characters.ValidateParentheses(str)
	str = "{{{{}]}}}"
	characters.ValidateParentheses(str)
	fmt.Printf("\n")

	str = "(((())))(((((())))))"
	characters.LongestValidParentheses(str)
	str = "(((()))))(())"
	characters.LongestValidParentheses(str)
	str = "(((())))(((())))"
	characters.LongestValidParentheses(str)
	fmt.Printf("\n")

	str = "Red rum sir is murder"
	characters.ValidPalindrome(str)
	str = "Red rum sir is not murder"
	characters.ValidPalindrome(str)
	str = ""
	characters.ValidPalindrome(str)
	fmt.Printf("\n")

	str = "qwertyghqydidapjdwoasodwja"
	characters.LongestSubstringWithoutRepeatingCharacters(str)
	str = "abcabcdbb"
	characters.LongestSubstringWithoutRepeatingCharacters(str)
	fmt.Printf("\n")

	arr = [10]int{1, 2, 3, 4, 5, 4, 3, 2, 1, 1}
	characters.FindPeakElement(arr[:])
	arr = [10]int{7, 2, 3, 4, 5, 4, 3, 2, 1, 6}
	characters.FindPeakElement(arr[:])
}
