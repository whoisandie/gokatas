package helpers

import (
	"strings"
)

// ReverseString : Reverse a string
func ReverseString(str string) string {
	strLen := len(str)
	reversed := make([]string, strLen)
	for index := range str {
		reversed[index] = string(str[strLen-index-1])
	}

	return strings.Join(reversed, "")
}
