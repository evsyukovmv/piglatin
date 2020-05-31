package piglatin

import (
	"testing"
)

var rules = map[string]map[string]string{
	"begin with consonant": {
		"pig":    "igpay",
		"latin":  "atinlay",
		"banana": "ananabay",
		"will":   "illway",
		"butler": "utlerbay",
		"happy":  "appyhay",
		"duck":   "uckday",
		"me":     "emay",
	},
	"begin with consonant clusters": {
		"smile":  "ilesmay",
		"string": "ingstray",
		"stupid": "upidstay",
		"glove":  "oveglay",
		"trash":  "ashtray",
		"floor":  "oorflay",
		"store":  "orestay",
	},
	"begin with vowel": {
		"eat":     "eatay",
		"omelet":  "omeletay",
		"are":     "areay",
		"egg":     "eggay",
		"explain": "explainay",
		"always":  "alwaysay",
		"ends":    "endsay",
		"I":       "Iay",
	},
}

var phrases = [...]map[string]string{
	{
		"phrase":   "Hello, World!",
		"expected": "Ellohay, Orldway!",
	},
	{
		"phrase":   "Pig Latin was widely used by The Three Stooges; for example, in the 1940 Hitler parody You Nazty Spy!",
		"expected": "Igpay Atinlay asway idelyway useday byay Ethay Eethray Oogesstay; orfay exampleay, inay ethay 1940 Itlerhay arodypay Ouyay Aztynay Spyay!",
	},
	{
		"phrase":   "",
		"expected": "",
	},
	{
		"phrase":   "!@#$%^&*()",
		"expected": "!@#$%^&*()",
	},
}

func TestConvertWord(t *testing.T) {
	for rule, items := range rules {
		for word, expected := range items {
			result := string(convertWord([]rune(word)))
			if result != expected {
				t.Errorf("test for Rule %q failed - results not match\nGot:\n%v\nExpected:\n%v", rule, result, expected)
				break
			}
		}
	}
}

func TestPigLatin(t *testing.T) {
	for _, phrase := range phrases {
		result := PigLatin(phrase["phrase"])
		expected := phrase["expected"]

		if result != expected {
			t.Errorf("test for Sentence failed - results not match\nGot:\n%v\nExpected:\n%v", result, expected)
			break
		}
	}
}
