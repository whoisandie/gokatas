package problems

// Abs : Absolute integer value
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}

// Som : Adds two numbers x and y
func Som(x, y int) int {
	return x + y
}

// Maxi : Maximum of x and y
func Maxi(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Mini : Minimum of x and y
func Mini(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Lcmu : LCM of absolute values of x and y
func Lcmu(x, y int) int {
	abx := Abs(x)
	aby := Abs(y)
	return abx * aby / Gcdi(abx, aby)
}

// Gcdi : GCD of absolute values of x and y
func Gcdi(x, y int) int {
	abx := Abs(x)
	aby := Abs(y)
	if aby == 0 {
		return abx
	}
	return Gcdi(aby, abx%aby)
}

// FParam Function definition type
type FParam func(int, int) int

// ReduceBySteps : Reduce By Steps function
func ReduceBySteps(f FParam, arr []int, init int) []int {
	lenArr := len(arr)
	r := make([]int, lenArr)
	for index, elem := range arr {
		var start int
		if index != 0 {
			start = r[index-1]
		} else {
			start = init
		}
		r[index] = f(start, elem)
	}
	return r
}
