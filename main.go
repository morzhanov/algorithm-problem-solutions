package main

import (
	"fmt"
	"math/rand"

	"github.com/morzhanov/algorithm-problem-solutions/array"
	"github.com/morzhanov/algorithm-problem-solutions/matrix"
)

func createArray() []int {
	values := make([]int, 20)
	for i := range values {
		values[i] = rand.Intn(100)
	}
	return values
}

func runMatrix() {
	mx := [][]rune{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
	}
	matrix.NumberOfIslands(mx)
	mx = [][]rune{
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
	}
	matrix.NumberOfIslands(mx)
	mx = [][]rune{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '1'},
	}
	matrix.NumberOfIslands(mx)
	fmt.Printf("\n")

	mx = [][]rune{
		{'6', '3', '9', '5', '7', '4', '1', '8', '2'},
		{'5', '4', '1', '8', '2', '9', '3', '7', '6'},
		{'7', '8', '2', '6', '1', '3', '9', '5', '4'},
		{'1', '9', '8', '4', '6', '7', '5', '2', '3'},
		{'3', '6', '5', '9', '8', '2', '4', '1', '7'},
		{'4', '2', '7', '1', '3', '5', '8', '6', '9'},
		{'9', '5', '6', '7', '4', '8', '2', '3', '1'},
		{'8', '1', '3', '2', '9', '6', '7', '4', '5'},
		{'2', '7', '4', '3', '5', '1', '6', '9', '8'},
	}
	matrix.ValidSudoku(mx)
	mx = [][]rune{
		{'6', '3', '9', '5', '7', '4', '1', '8', '2'},
		{'5', '4', '1', '8', '2', '9', '3', '7', '6'},
		{'7', '8', '2', '6', '1', '3', '.', '.', '.'},
		{'1', '9', '8', '.', '6', '7', '5', '2', '3'},
		{'3', '6', '5', '9', '8', '2', '4', '1', '7'},
		{'4', '2', '7', '1', '.', '5', '8', '6', '9'},
		{'9', '5', '6', '7', '4', '8', '2', '3', '1'},
		{'8', '1', '3', '2', '9', '.', '7', '4', '5'},
		{'.', '.', '.', '3', '5', '1', '6', '9', '8'},
	}
	matrix.ValidSudoku(mx)
	mx = [][]rune{
		{'6', '3', '9', '5', '7', '4', '1', '8', '2'},
		{'5', '4', '1', '8', '2', '9', '3', '7', '6'},
		{'7', '8', '2', '6', '1', '3', '.', '.', '.'},
		{'1', '.', '8', '.', '6', '7', '5', '2', '3'},
		{'3', '6', '5', '9', '8', '2', '4', '1', '7'},
		{'4', '2', '7', '1', '.', '5', '8', '6', '9'},
		{'9', '5', '6', '7', '4', '8', '2', '3', '1'},
		{'8', '1', '3', '2', '9', '.', '7', '4', '5'},
		{'.', '9', '.', '3', '5', '1', '6', '.', '8'},
	}
	matrix.ValidSudoku(mx)
	fmt.Printf("\n")
}

func runArray() {
	arr := [10]int{1, 2, 4, 5, 2, 3, 5, 1, 2, 4}
	array.FindPairWithGivenSumInArray(arr[:], 8)
	fmt.Printf("\n")

	arr = [10]int{1, -4, 4, 2, 2, -2, 4, 5, -8, -4}
	array.FindAllSubarraysWithZeroSum(arr[:])
	fmt.Printf("\n")

	arr = [10]int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9}
	array.FindDuplicateElementInLimitedRangeArray(arr[:])
	fmt.Printf("\n")

	str := "(){{{{}}}}"
	array.ValidateParentheses(str)
	str = "{{{({}){([()])}}}}"
	array.ValidateParentheses(str)
	str = "{{{{}]}}}"
	array.ValidateParentheses(str)
	fmt.Printf("\n")

	str = "(((())))(((((())))))"
	array.LongestValidParentheses(str)
	str = "(((()))))(())"
	array.LongestValidParentheses(str)
	str = "(((())))(((())))"
	array.LongestValidParentheses(str)
	fmt.Printf("\n")

	str = "Red rum sir is murder"
	array.ValidPalindrome(str)
	str = "Red rum sir is not murder"
	array.ValidPalindrome(str)
	str = ""
	array.ValidPalindrome(str)
	fmt.Printf("\n")

	str = "qwertyghqydidapjdwoasodwja"
	array.LongestSubstringWithoutRepeatingarray(str)
	str = "abcabcdbb"
	array.LongestSubstringWithoutRepeatingarray(str)
	fmt.Printf("\n")

	arr = [10]int{1, 2, 3, 4, 5, 4, 3, 2, 1, 1}
	array.FindPeakElement(arr[:])
	arr = [10]int{7, 2, 3, 4, 5, 4, 3, 2, 1, 6}
	array.FindPeakElement(arr[:])
	fmt.Printf("\n")

	array.GuessNumberHigherOrLower(10, 8)
	array.GuessNumberHigherOrLower(100, 43)
	array.GuessNumberHigherOrLower(500, 222)
	array.GuessNumberHigherOrLower(1000, 998)
	fmt.Printf("\n")

	str = "PAYPALISHIRING"
	array.ZigZagConversion(str, 2)
	array.ZigZagConversion(str, 3)
	array.ZigZagConversion(str, 4)
	array.ZigZagConversion(str, 5)
	fmt.Printf("\n")

	secret := "1234"
	_guess := "4321"
	array.BullsAndCows(secret, _guess)
	secret = "4567"
	_guess = "4576"
	array.BullsAndCows(secret, _guess)
	secret = "8762"
	_guess = "1347"
	array.BullsAndCows(secret, _guess)
	secret = "1234"
	_guess = "5678"
	array.BullsAndCows(secret, _guess)
	secret = "1234"
	_guess = "1234"
	array.BullsAndCows(secret, _guess)
	secret = "1122"
	_guess = "2211"
	array.BullsAndCows(secret, _guess)
	secret = "1212123"
	_guess = "2221113"
	array.BullsAndCows(secret, _guess)
	fmt.Printf("\n")

	strs := []string{"abc", "abcd", "abcde", "abcdef"}
	array.LongestCommonPrefix(strs)
	strs = []string{"abc", "abcd", "abcde", "abcdef", "a"}
	array.LongestCommonPrefix(strs)
	strs = []string{"abc", "abcd", "abcde", "abcdef", "k"}
	array.LongestCommonPrefix(strs)
	fmt.Printf("\n")

	ratings := []int{1, 4, 3, 3, 3, 1}
	array.Candy(ratings)
	ratings = []int{1, 2, 3, 2, 1, 1}
	array.Candy(ratings)
	ratings = []int{1, 1, 1, 1, 1, 1}
	array.Candy(ratings)
	ratings = []int{1, 1, 2, 1, 1, 2}
	array.Candy(ratings)
	ratings = []int{3, 1, 2, 1, 1, 3}
	array.Candy(ratings)
	fmt.Printf("\n")

	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	array.TrappingRainWatter(heights)
	fmt.Printf("\n")

	arr2 := []int{1, 2, 3}
	res := array.Permutations(arr2[:], 0)
	fmt.Printf("Permutations: %v\n", res)
	arr2 = []int{1, 2, 3, 4}
	res = array.Permutations(arr2[:], 0)
	fmt.Printf("Permutations: %v\n", res)
	arr2 = []int{1, 2, 3, 4, 5}
	res = array.Permutations(arr2[:], 0)
	fmt.Printf("Permutations: %v\n", res)
	fmt.Printf("\n")

	array.HappyNumber(4)
	array.HappyNumber(18)
	array.HappyNumber(19)
	array.HappyNumber(23)
	array.HappyNumber(761)
	fmt.Printf("\n")
}

func main() {
	runMatrix()
	runArray()
}
