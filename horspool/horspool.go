package horspool

func createShiftTable(pattern string) map[byte]int {
	shiftTable := make(map[byte]int)

	for i := 0; i < len(pattern)-1; i++ {
		index := pattern[i]
		shiftTable[index] = len(pattern) - i - 1
	}
	return shiftTable
}

func createReverseShiftTable(pattern string) map[byte]int {
	shiftTable := make(map[byte]int)

	for i := 1; i < len(pattern); i++ {
		index := pattern[i]
		if _, isExist := shiftTable[index]; !isExist {
			shiftTable[index] = i
		}
	}
	return shiftTable
}

func calculateShiftAmount(shiftTable map[byte]int, char byte, patternLength int) int {
	if shiftAmount, isExists := shiftTable[char]; isExists {
		return shiftAmount
	}
	return patternLength

}

// Find finds the index of first matching pattern in text using horspool algorithm
// returns -1 if pattern does not exist in text
func Find(text string, pattern string) int {
	shiftTable := createShiftTable(pattern)

	textLength := len(text)
	patternLength := len(pattern)

	needle := patternLength

	for needle <= textLength {
		textSlice := text[needle-patternLength : needle]
		if textSlice != pattern {
			needle += calculateShiftAmount(shiftTable, text[needle-1], patternLength)
		} else {
			return needle - patternLength
		}
	}

	return -1
}

// FindLast finds the index of last matching pattern in thex using horspool algorithm
// returns -1 if pattern does not exist in text
func FindLast(text string, pattern string) int {
	shiftTable := createReverseShiftTable(pattern)

	textLength := len(text)
	patternLength := len(pattern)

	needle := textLength - patternLength

	for needle >= 0 {
		textSlice := text[needle : needle+patternLength]
		if textSlice != pattern {
			needle -= calculateShiftAmount(shiftTable, text[needle], patternLength)
		} else {
			return needle
		}
	}

	return -1
}
