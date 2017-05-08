package helpers

// Combinations : generates combinations of numbers
func Combinations(list []int, k int) [][]int {
	listLen := len(list)

	if k > len(list) || k <= 0 {
		return [][]int{}
	}

	if k == len(list) {
		return [][]int{list}
	}

	if k == 1 {
		tCombs := [][]int{}
		for i := 0; i < listLen; i++ {
			temp := []int{list[i]}
			tCombs = append(tCombs, temp)
		}
		return tCombs
	}

	combs := [][]int{}
	for i := 0; i < listLen-k+1; i++ {
		head := list[i : i+1]
		tail := Combinations(list[i+1:], k-1)

		for j := 0; j < len(tail); j++ {
			temp := make([]int, 0)
			temp = append(temp, head...)
			temp = append(temp, tail[j]...)
			combs = append(combs, temp)
		}
	}

	return combs
}
