package piglatin

import "unicode"

var (
	vowels = map[rune]struct{}{
		'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {}, 'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
	}
)

func capitalize(runes []rune) {
	for i := range runes {
		runes[i] = unicode.ToLower(runes[i])
	}
	runes[0] = unicode.ToUpper(runes[0])
}

func isLatinLetter(r rune) bool {
	return unicode.IsLetter(r) && r <= unicode.MaxASCII
}

func isVowel(r rune) bool {
	_, ok := vowels[r]
	return ok
}
