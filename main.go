package main

import (
	"math/rand"
)

func createArray() []int {
	values := make([]int, 20)
	for i := range values {
		values[i] = rand.Intn(100)
	}
	return values
}

func main() {
	arr := [10]int{1, 2, 4, 5, 2, 3, 5, 1, 2, 4}
	FindPairWithGivenSumInArray(arr[:], 8)
}
