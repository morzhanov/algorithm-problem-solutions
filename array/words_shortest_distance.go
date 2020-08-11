package array

import (
	"fmt"
	"math"
)

// WordsShortestDistance test
// Given a list of words and two words word1 and word2
// return the shortest distance between these two words in the list.
func WordsShortestDistance(words []string, word1 string, word2 string) {
	first := -1
	last := -1
	min := len(words)

	for i := 0; i < len(words); i++ {
		if words[i] == word1 {
			first = i
			if last != -1 {
				curr := int(math.Abs(float64(last)-float64(first))) - 1
				if min > curr {
					min = curr
				}
			}
		} else if words[i] == word2 {
			last = i
			if first != -1 {
				curr := int(math.Abs(float64(last)-float64(first))) - 1
				if min > curr {
					min = curr
				}
			}
		}
	}

	fmt.Printf("Shortest distance between words %v and %v is %v\n", word1, word2, min)
}

// WordsShortestDistanceN2 test N^2 complexity
func WordsShortestDistanceN2(words []string, word1 string, word2 string) {
	pos1 := make([]int, 0)
	pos2 := make([]int, 0)

	for i := 0; i < len(words); i++ {
		if words[i] == word1 {
			pos1 = append(pos1, i)
		}
		if words[i] == word2 {
			pos2 = append(pos2, i)
		}
	}

	min := len(words)
	for i := 0; i < len(pos1); i++ {
		for j := 0; j < len(pos2); j++ {
			abs := int(math.Abs(float64(pos1[i])-float64(pos2[j]))) - 1
			if abs < min {
				min = abs
			}
		}
	}
	fmt.Printf("Shortest distance between words %v and %v is %v\n", word1, word2, min)
}
