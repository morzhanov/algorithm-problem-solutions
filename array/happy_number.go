package array

import "fmt"

// HappyNumber test
// Write an algorithm to determine if a number is "happy".
// What is an happy number can be shown in the following example:
// 19 is a happy number
// 1^2 + 9^2 = 82
// 8^2 + 2^2 = 68
// 6^2 + 8^2 = 100
// 1^2 + 0^2 + 0^2 = 1
func HappyNumber(number int) {
	initialNumber := number
	s := map[int]bool{}

	for ok := false; ok == false; _, ok = s[number] {
		s[number] = true
		number = sum(number)

		if number == 1 {
			fmt.Printf("%v is a happy number!\n", initialNumber)
			return
		}
	}
	fmt.Printf("%v is not a happy number\n", initialNumber)
}

func sum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	fmt.Printf("HappyNumber current sum: %v\n", sum)
	return sum
}
