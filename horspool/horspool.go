package horspool

import ()

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
	} else {
		return patternLength
	}
}

// Find finds the index of first matching pattern in text using horspool algorithm
// returns -1 if pattern does not exist in text
func Find(text string, pattern string) int {
	shiftTable := createShiftTable(pattern)

	textLength := len(text)
	patternLength := len(pattern)

	needle := patternLength - 1
	textCompareIndex := patternLength - 1
	patternCompareIndex := patternLength - 1

	matchCount := 0

	for {
		if textCompareIndex > textLength {
			return -1
		}

		textChar := text[textCompareIndex]
		if textChar != pattern[patternCompareIndex] {
			needle = textCompareIndex + calculateShiftAmount(shiftTable, textChar, patternLength)
			textCompareIndex = needle
			patternCompareIndex = patternLength - 1
			matchCount = 0
		} else {
			matchCount++
			if matchCount == patternLength {
				return needle - patternLength + 1
			}
			patternCompareIndex--
			textCompareIndex--
		}
	}
}

// FindLast finds the index of last matching pattern in thex using horspool algorithm
// returns -1 if pattern does not exist in text
func FindLast(text string, pattern string) int {
	shiftTable := createReverseShiftTable(pattern)

	textLength := len(text)
	patternLength := len(pattern)

	needle := textLength - patternLength
	textCompareIndex := needle
	patternCompareIndex := 0
	matchCount := 0

	for {
		if textCompareIndex < 0 {
			return -1
		}

		textChar := text[textCompareIndex]
		if textChar != pattern[patternCompareIndex] {
			needle = textCompareIndex - calculateShiftAmount(shiftTable, textChar, patternLength)
			textCompareIndex = needle
			patternCompareIndex = 0
			matchCount = 0
		} else {
			matchCount++
			if matchCount == patternLength {
				return needle
			}
			patternCompareIndex++
			textCompareIndex++
		}
	}
}
