package characters

import (
	"fmt"

	"github.com/morzhanov/algorithm-problem-solutions/utils"
)

// ZigZagConversion test
// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
//  (you may want to display this pattern in a fixed font for better legibility)
// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
func ZigZagConversion(str string, numRows int) {
	res := make([][]rune, numRows)
	arr := []rune(str)
	n := len(arr)
	direction := 0 // 0 - down, 1 - up
	row := 0
	maxRowLen := 0

	for i := 0; i < n; i++ {
		res[row] = append(res[row], arr[i])

		if row == 0 && direction == 1 {
			direction = 0
			row++
		} else if row == numRows-1 && direction == 0 {
			direction = 1
			row--
		} else if direction == 0 {
			row++
		} else {
			row--
		}
		rowLen := len(res[row])
		if rowLen > maxRowLen {
			maxRowLen = rowLen
		}
	}

	fmt.Printf("%v in ZigZag:\n", str)
	utils.PrintMatrixInZigZag(res, maxRowLen)
}
