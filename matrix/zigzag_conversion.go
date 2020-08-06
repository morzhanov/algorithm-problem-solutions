package matrix

import (
	"fmt"

	"github.com/morzhanov/algorithm-problem-solutions/utils"
)

// ZigZagConcersion test
// Given a 2-d grid map of '1's (land) and '0's (water), count the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.
func ZigZagConcersion(mx [][]rune) {
	if len(mx[0]) == 0 || len(mx) == 0 {
		fmt.Printf("Count of islands: 0\n")
		return
	}

	n := len(mx)
	m := len(mx[0])
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mx[i][j] == '1' {
				count++
				merge(mx, i, j)
				fmt.Printf("Matrix after merge phase:\n")
				utils.PrintMatrix(mx)
			}
		}
	}

	fmt.Printf("Count of islands: %v\n", count)
}

func merge(mx [][]rune, i int, j int) {
	n := len(mx)
	m := len(mx[0])

	if i < 0 || i >= n || j < 0 || j >= m || mx[i][j] != '1' {
		return
	}

	mx[i][j] = 'X'
	merge(mx, i-1, j)
	merge(mx, i+1, j)
	merge(mx, i, j-1)
	merge(mx, i, j+1)
}
