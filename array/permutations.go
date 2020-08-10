package array

// Permutations test
// Given a collection of numbers, return all possible permutations
func Permutations(arr []int, start int) [][]int {
	temp := make([]int, 0)
	temp = append(temp, arr[start:]...)
	res := make([][]int, 0)

	if len(temp) == 2 {
		res = append(res, []int{temp[0], temp[1]})
		res = append(res, []int{temp[1], temp[0]})
		return res
	}

	for i := 0; i < len(temp); i++ {
		temp[0], temp[i] = arr[i], arr[0]
		children := Permutations(arr, start+1)

		for j := 0; j < len(children); j++ {
			curr := make([]int, 1)
			curr[0] = temp[0]
			curr = append(curr, children[j]...)
			res = append(res, curr)
		}
	}
	return res
}
