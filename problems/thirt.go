package problems

import (
	"strconv"

	"github.com/whoisandie/gokatas/helpers"
)

// Thirt : Compute thirt
func Thirt(number int) int {
	var sequence = [6]int{1, 10, 9, 12, 3, 4}
	numberString := strconv.FormatInt(int64(number), 10)
	reversedNumberString := helpers.ReverseString(numberString)
	digitSum := 0

	for index, item := range reversedNumberString {
		seqLen := len(sequence)
		currentNumber, _ := strconv.Atoi(string(item))
		currIndex := index
		if index >= seqLen {
			currIndex = index % seqLen
		}
		digitSum += sequence[currIndex] * currentNumber
	}

	if number-digitSum != 0 {
		return Thirt(digitSum)
	}

	return number
}
