package problems

// TortoiseRace : race
func TortoiseRace(v1, v2, g int) [3]int {
	result := [3]int{-1, -1, -1}

	if v1 > v2 {
		return result
	}

	timeToCatch := float64(g) / (float64(v2) - float64(v1)) * 3600
	result[0] = int(timeToCatch) / 3600
	timeToCatch = float64(int(timeToCatch) % 3600)
	result[1] = int(timeToCatch / 60.0)
	result[2] = int(timeToCatch) % 60

	return result
}
