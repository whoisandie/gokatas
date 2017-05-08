package problems

// RemovNb : Cheating friend
func RemovNb(m uint64) [][2]uint64 {
	n := m
	sum := uint64(n * (n + 1) / 2)
	var res [][2]uint64
	for a := uint64(1); a <= n; a++ {
		if (sum-a)%(a+1) == 0 {
			b := (sum - a) / (a + 1)
			if b < n {
				res = append(res, [2]uint64{a, b})
			}
		}
	}
	return res
}

// func RemovNb(number uint64) [][2]uint64 {
// 	sum := ((number * (number + 1)) / 2)
// 	list := make([]int, number)
// 	for index := range list {
// 		list[index] = index + 1
// 	}
// 	combs := examples.Combinations(list, 2)

// 	var result [][2]uint64
// 	for _, item := range combs {
// 		a, b := uint64(item[0]), uint64(item[1])
// 		if uint64(a*b+a+b) == sum {
// 			result = append(result, [2]uint64{a, b})
// 			result = append(result, [2]uint64{b, a})
// 		}
// 	}
// 	return result
// }
