package main

// Pos returns position of rune in the run array
func Pos(arr []rune, val rune) int {
	for p, v := range arr {
		if v == val {
			return p
		}
	}
	return -1
}
