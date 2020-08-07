package main

import (
	"fmt"
	"math/rand"

	"github.com/morzhanov/algorithm-problem-solutions/characters"
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

func runCharacters() {
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
	fmt.Printf("\n")

	characters.GuessNumberHigherOrLower(10, 8)
	characters.GuessNumberHigherOrLower(100, 43)
	characters.GuessNumberHigherOrLower(500, 222)
	characters.GuessNumberHigherOrLower(1000, 998)
	fmt.Printf("\n")

	str = "PAYPALISHIRING"
	characters.ZigZagConversion(str, 2)
	characters.ZigZagConversion(str, 3)
	characters.ZigZagConversion(str, 4)
	characters.ZigZagConversion(str, 5)
	fmt.Printf("\n")

	secret := "1234"
	_guess := "4321"
	characters.BullsAndCows(secret, _guess)
	secret = "4567"
	_guess = "4576"
	characters.BullsAndCows(secret, _guess)
	secret = "8762"
	_guess = "1347"
	characters.BullsAndCows(secret, _guess)
	secret = "1234"
	_guess = "5678"
	characters.BullsAndCows(secret, _guess)
	secret = "1234"
	_guess = "1234"
	characters.BullsAndCows(secret, _guess)
	secret = "1122"
	_guess = "2211"
	characters.BullsAndCows(secret, _guess)
	secret = "1212123"
	_guess = "2221113"
	characters.BullsAndCows(secret, _guess)
	fmt.Printf("\n")

	strs := []string{"abc", "abcd", "abcde", "abcdef"}
	characters.LongestCommonPrefix(strs)
	strs = []string{"abc", "abcd", "abcde", "abcdef", "a"}
	characters.LongestCommonPrefix(strs)
	strs = []string{"abc", "abcd", "abcde", "abcdef", "k"}
	characters.LongestCommonPrefix(strs)
	fmt.Printf("\n")
}

func main() {
	runMatrix()
	runCharacters()
}
