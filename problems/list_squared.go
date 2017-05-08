package problems

import (
	"math"
)

// SumFactors : sum of factors of an integer
func SumFactors(n int) int {
	sum := 0
	start := 2
	limit := n
	for limit > 0 {
		if limit == 1 {
			sum++
			break
		}

		if n%start == 0 {
			sum += (limit * limit)
			limit = n / start
		}
		start++
	}
	return sum
}

// ListSquared : Number whose squared divisors is a square
func ListSquared(m, n int) [][]int {
	list := [][]int{}
	for i := m; i <= n; i++ {
		sum := SumFactors(i)
		sqrtDiff := math.Mod(math.Sqrt(float64(sum)), 1)
		if sqrtDiff == 0 {
			list = append(list, []int{i, sum})
		}
	}

	return list
}
