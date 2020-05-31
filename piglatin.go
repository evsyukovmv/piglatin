package piglatin

import (
	"unicode"
)

const (
	ending = "ay"
)

// PigLatin converts words in English to Pig Latin (https://en.wikipedia.org/wiki/Pig_Latin)
func PigLatin(s string) string {
	runes := []rune(s)
	result := make([]rune, 0, len(runes)*2)

	position := 0
	for i, r := range runes {
		if isLatinLetter(r) {
			continue
		}
		if i > position {
			result = append(result, convertWord(runes[position:i])...)
		}
		result = append(result, r)
		position = i + 1
	}
	return string(result)
}

func convertWord(runes []rune) []rune {
	result := make([]rune, len(runes), len(runes)+len(ending))
	copy(result, runes)

	if !isVowel(result[0]) {
		result = moveConsonants(result)
	}
	result = addVowels(result)

	if unicode.IsUpper(runes[0]) {
		capitalize(result)
	}

	return result
}

func addVowels(runes []rune) []rune {
	return append(runes, []rune(ending)...)
}

func moveConsonants(runes []rune) []rune {
	consonantsEnd := 0
	for i, r := range runes {
		if isVowel(r) {
			consonantsEnd = i
			break
		}
	}

	return append(runes[consonantsEnd:], runes[:consonantsEnd]...)
}
