package main

func canBeTypedWords(text string, brokenLetters string) int {
	broken := make(map[rune]bool)
	for _, b := range brokenLetters {
		broken[b] = true
	}

	result := 0
	canType := 1
	for _, char := range text {
		if char == ' ' {
			result += canType
			canType = 1
		} else if canType == 1 && broken[char] {
			canType = 0
		}
	}
	result += canType

	return result
}
