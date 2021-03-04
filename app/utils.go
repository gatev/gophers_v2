package main

import (
	"strings"
)

const G = "g"
const GE = "ge"
const XR = "xr"
const Q = "q"
const U = "u"
const OGO = "ogo"

var translations map[string]string

func translate(word string) string {
	var result string
	var seqOfConsonant string
	var endIndexOfConsonant int
	r := []rune(strings.ToLower(word))
	if isVowel(r[0]) {
		result = G + word
	} else if len(word) > 1 && strings.ToLower(word)[0:2] == XR {
		result = GE + word
	} else if !isVowel(r[0]) {
		endIndexOfConsonant = countSeqOfConsonant(r)
		seqOfConsonant = word[:endIndexOfConsonant]

		if seqOfConsonant[len(seqOfConsonant)-1:] == Q && word[endIndexOfConsonant:endIndexOfConsonant+1] == U {
			seqOfConsonant = seqOfConsonant + U
			result = word[endIndexOfConsonant+1:] + seqOfConsonant + OGO
		} else {
			result = word[endIndexOfConsonant:] + seqOfConsonant + OGO
		}
	}
	return result
}

func isVowel(char rune) bool {
	var result bool
	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'y':
		result = true
	}
	return result
}

func countSeqOfConsonant(word []rune) int {
	var result int = 1

	for i := 1; i < len(word); i++ {
		if !isVowel(word[i]) {
			result++
		} else {
			break
		}
	}
	return result
}