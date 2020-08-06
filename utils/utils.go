package utils

import "fmt"

// Pos returns position of rune in the run array
func Pos(arr []rune, val rune) int {
	for p, v := range arr {
		if v == val {
			return p
		}
	}
	return -1
}

// PrintMatrix function prints matrix of runes
func PrintMatrix(arr [][]rune) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v\n", string(arr[i]))

	}
}
