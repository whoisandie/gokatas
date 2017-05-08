package problems

import "strings"

// ToNato : Convert sentence to nato
func ToNato(words string) string {
	nato := []string{
		"Alfa",
		"Bravo",
		"Charlie",
		"Delta",
		"Echo",
		"Foxtrot",
		"Golf",
		"Hotel",
		"India",
		"Juliett",
		"Kilo",
		"Lima",
		"Mike",
		"November",
		"Oscar",
		"Papa",
		"Quebec",
		"Romeo",
		"Sierra",
		"Tango",
		"Uniform",
		"Victor",
		"Whiskey",
		"X-ray",
		"Yankee",
		"Zulu"}
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	natoMap := map[string]string{}
	var result string
	for index, char := range alphabets {
		natoMap[string(char)] = nato[index]
	}

	lowercaseWordsNoSpaces := strings.Join(strings.Split(strings.ToLower(words), " "), "")
	for index, char := range lowercaseWordsNoSpaces {
		result += natoMap[string(char)]
		if index != len(words)-1 {
			result += " "
		}
	}

	return result
}
