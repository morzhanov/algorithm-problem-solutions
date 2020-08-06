package matrix

import (
	"fmt"

	"github.com/morzhanov/algorithm-problem-solutions/utils"
)

// ValidSudoku test
// Determine if a Sudoku is valid.
// The Sudoku board could be partially filled, where empty cells are filled with the character '.'
func ValidSudoku(mx [][]rune) {
	if len(mx) != 9 || len(mx[0]) != 9 {
		fmt.Printf("Matrix is not valid sudoku board\n")
		utils.PrintMatrix(mx)
	}

	for i := 0; i < 9; i++ {
		arr := make([]rune, 9)
		for j := 0; j < 9; j++ {
			el := mx[i][j]
			if el != '.' {
				if utils.Pos(arr, el) != -1 {
					fmt.Printf("Matrix is not valid sudoku board (ROW check phase)\n")
					fmt.Printf("ROW = %v, COLUMN = %v, VALUE = %v\n", i, j, el)
					return
				}
				arr = append(arr, el)
			}
		}
	}

	for i := 0; i < 9; i++ {
		arr := make([]rune, 9)
		for j := 0; j < 9; j++ {
			el := mx[j][i]
			if el != '.' {
				if utils.Pos(arr, el) != -1 {
					fmt.Printf("Matrix is not valid sudoku board (COLUMN check phase)\n")
					fmt.Printf("ROW = %v, COLUMN = %v, VALUE = %v\n", i, j, el)
					return
				}
				arr = append(arr, el)
			}
		}
	}

	for block := 0; block < 9; block++ {
		arr := make([]rune, 9)
		for i := block / 3 * 3; i < block/3*3+3; i++ {
			for j := block % 3 * 3; j < block%3*3+3; j++ {
				el := mx[i][j]
				if el != '.' {
					if utils.Pos(arr, el) != -1 {
						fmt.Printf("Matrix is not valid sudoku board (BLOCK check phase)\n")
						fmt.Printf("BLOCK = %v, ROW = %v, COLUMN = %v, VALUE = %v\n", block, i, j, el)
						return
					}
					arr = append(arr, el)
				}
			}
		}
	}

	fmt.Printf("Matrix is a valid sudoku board\n")
}
