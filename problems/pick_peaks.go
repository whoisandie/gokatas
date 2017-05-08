package problems

// PosPeaks struct
type PosPeaks struct {
	Pos   []int
	Peaks []int
}

// PickPeaks : Pick peaks in a slice
func PickPeaks(array []int) PosPeaks {
	arrLen := len(array)
	peaks := PosPeaks{Pos: []int{}, Peaks: []int{}}
	platStart := false
	platEnd := false
	var platIdx int
	for i := 1; i < arrLen-1; i++ {
		prevDiff := array[i] - array[i-1]
		currDiff := array[i+1] - array[i]

		if prevDiff > 0 && currDiff == 0 {
			platIdx = i
			platStart = true
		}

		if prevDiff == 0 && currDiff < 0 {
			platEnd = true
		}

		if platStart && platEnd && array[platIdx] > array[i+1] {
			peaks.Pos = append(peaks.Pos, platIdx)
			peaks.Peaks = append(peaks.Peaks, array[platIdx])
			platStart = false
			platEnd = false
		}

		if prevDiff > 0 && currDiff < 0 {
			peaks.Pos = append(peaks.Pos, i)
			peaks.Peaks = append(peaks.Peaks, array[i])
		}
	}

	return peaks
}
