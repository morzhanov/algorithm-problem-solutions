package array

import (
	"fmt"
)

var number = 0
var iterations = 0

// GuessNumberHigherOrLower test
// We are playing the Guess Game. The game is as follows:
// I pick a number from 1 to n. You have to guess which number I picked.
// Every time you guess wrong, I'll tell you whether the number is higher or lower.
// You call a pre-defined API guess(int num) which returns 3 possible results (-1, 1, or 0):
// -1 : My number is lower
//  1 : My number is higher
//  0 : Congrats! You got it!
// Example: n = 10, I pick 6. Return 6.
func GuessNumberHigherOrLower(n int, val int) {
	number = val
	iterations = 0
	res := guessNumber(0, n)
	fmt.Printf("Found value %v, iteration taken %v\n", res, iterations)
}

func guessNumber(l int, h int) int {
	iterations++
	med := (h + l) / 2
	res := guess(med)
	if res == -1 {
		return guessNumber(med, h)
	}
	if res == 1 {
		return guessNumber(0, med-1)
	}
	return med
}

func guess(val int) int {
	if val > number {
		return 1
	}
	if val < number {
		return -1
	}
	return 0
}
