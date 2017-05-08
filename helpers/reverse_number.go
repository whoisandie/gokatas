package helpers

// ReverseNumber : Reverse a number
func ReverseNumber(number int) int {
	temp := 0
	for number > 0 {
		remainder := number % 10
		temp *= 10
		temp += remainder
		number /= 10
	}

	return temp
}
