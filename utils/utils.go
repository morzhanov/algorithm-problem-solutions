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

// PrintMatrixInZigZag function prints matrix of runes in zig-zag style
func PrintMatrixInZigZag(arr [][]rune, maxRowLen int) {
	for i := 0; i < len(arr); i++ {
		delimiter := maxRowLen/len(arr[i]) + 1
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v", string(arr[i][j]))
			for k := 0; k < delimiter; k++ {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

// Min function returns min int
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
