package problems

// HasUniqueChar :finds unique character in string
func HasUniqueChar(str string) bool {
	charMap := map[string]string{}
	for _, elem := range str {
		char := string(elem)
		_, charPresent := charMap[char]
		if !charPresent {
			charMap[char] = char
		} else {
			return false
		}

	}
	return true
}
