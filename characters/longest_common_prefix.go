package characters

import (
	"fmt"
)

// LongestCommonPrefix test
// Write a function to find the longest common prefix string amongst an array of strings.
func LongestCommonPrefix(str []string) {
	arr := make([][]rune, len(str))
	prefix := []rune(str[0])
	for i := 0; i < len(str); i++ {
		arr[i] = []rune(str[i])
		if len(prefix) > len(arr[i]) {
			prefix = arr[i]
		}
	}

	for i := 0; i < len(str); i++ {
		curr := make([]rune, 0)
		for j := 0; j < len(prefix); j++ {
			if arr[i][j] != prefix[j] {
				prefix = curr
				break
			}
			curr = append(curr, prefix[j])
		}
	}

	fmt.Printf("Longest common prefix: %v\n", string(prefix))
}
