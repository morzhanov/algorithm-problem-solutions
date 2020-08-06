package characters

import (
	"fmt"
	"strings"

	"github.com/morzhanov/algorithm-problem-solutions/utils"
)

// LongestSubstringWithoutRepeatingCharacters test
// given a string, find the length of the longest substring without repeating characters
func LongestSubstringWithoutRepeatingCharacters(str string) {
	if str == "" {
		fmt.Printf("\"%v\" is empty\n", str)
		return
	}

	str = strings.ToLower(str)
	str = strings.Replace(str, " ", "", -1)
	arr := []rune(str)
	n := len(arr)
	window := make([]rune, 0)
	start := 0
	res := 0
	var resChars []rune

	for i := 0; i < n; i++ {
		pos := utils.Pos(window, arr[i])
		if pos != -1 {
			wLen := len(window)
			if res < wLen {
				res = wLen
				resChars = window
			}
			for k := 0; k < wLen && start+wLen <= n; k++ {
				if window[k] == arr[i] {
					start = start + k + 1
					window = arr[start : start+wLen]
					break
				}
			}
		} else {
			window = append(window, arr[i])
		}
	}

	fmt.Printf("Longest substring without repeating chars: %v\n", res)
	fmt.Printf("Sequence: %v\n", string(resChars))
}
