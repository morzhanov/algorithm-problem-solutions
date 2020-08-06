package matrix

import (
	"fmt"

	"github.com/morzhanov/algorithm-problem-solutions/utils"
)

// NumberOfIslands test
// Given a 2-d grid map of '1's (land) and '0's (water), count the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.
func NumberOfIslands(arr [][]rune) {
	if len(arr[0]) == 0 || len(arr) == 0 {
		fmt.Printf("Count of islands: 0\n")
		return
	}

	n := len(arr)
	m := len(arr[0])
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == '1' {
				count++
				merge(arr, i, j)
				fmt.Printf("Matrix after merge phase:\n")
				utils.PrintMatrix(arr)
			}
		}
	}

	fmt.Printf("Count of islands: %v\n", count)
}

func merge(arr [][]rune, i int, j int) {
	n := len(arr)
	m := len(arr[0])

	if i < 0 || i >= n || j < 0 || j >= m || arr[i][j] != '1' {
		return
	}

	arr[i][j] = 'X'
	merge(arr, i-1, j)
	merge(arr, i+1, j)
	merge(arr, i, j-1)
	merge(arr, i, j+1)
}
