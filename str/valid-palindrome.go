package str

import (
	"fmt"
	"strings"
)

// ValidPalindrome test
// check is string valid palindrome
// empty string is valid
func ValidPalindrome(str string) {
	if str == "" {
		fmt.Printf("\"%v\" is a valid palindrome\n", str)
		return
	}

	initialString := str
	str = strings.ToLower(str)
	str = strings.Replace(str, " ", "", -1)
	arr := []rune(str)

	idx := 0
	for idx != len(str) {
		if arr[idx] != arr[len(arr)-idx-1] {
			fmt.Printf("\"%v\" is not a valid palindrome\n", initialString)
			return
		}
		idx++
	}

	fmt.Printf("\"%v\" is a valid palindrome\n", initialString)
}
