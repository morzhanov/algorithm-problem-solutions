package characters

import (
	"fmt"

	"github.com/morzhanov/algorithm-problem-solutions/utils"
)

// BullsAndCows test
// You are playing the following Bulls and Cows game with your friend:
// You write down a number and ask your friend to guess what the number is.
// Each time your friend makes a guess, you provide a hint that indicates
// how many digits in said guess match your secret number exactly in both digit
// and position (called "bulls") and how many digits match the secret number
// but locate in the wrong position (called "cows").
// Your friend will use successive guesses and hints to eventually derive the secret number.
//
// For example:
// Secret number: "1807"
// Friend's guess: "7810"
// Hint: 1 bull and 3 cows. (The bull is 8, the cows are 0, 1 and 7.)
// Write a function to return a hint according to the secret number and friend's guess,
// use B to indicate the bulls and C to indicate the cows.
// In the above example, your function should return "1B3C".
func BullsAndCows(secret string, guess string) {
	b := 0
	c := 0
	secretArr := []rune(secret)
	guessArr := []rune(guess)

	for i := 0; i < len(secret); i++ {
		current := guessArr[i]
		pos := utils.Pos(secretArr, current)
		if guessArr[i] == secretArr[i] {
			b++
		} else if pos != -1 {
			c++
		}
	}
	fmt.Printf("%vB%vC\n", b, c)
}
