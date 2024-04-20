package german

import (
	"strings"

	"github.com/kljensen/snowball/romance"
	"github.com/kljensen/snowball/snowballword"
)

func IsStopWord(word string) bool {

	return false
}

func isLowerVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'y', 'ä', 'ö', 'ü':
		return true
	}
	return false
}

func isValidSEnding(r rune) bool {
	switch r {
	case 'b', 'd', 'f', 'g', 'h', 'k', 'l', 'm', 'n', 'r', 't':
		return true
	}
	return false
}

func isValidStEnding(r rune) bool {
	switch r {
	case 'b', 'd', 'f', 'g', 'h', 'k', 'l', 'm', 'n', 't':
		return true
	}
	return false
}

func r1r2(word *snowballword.SnowballWord) (r1start int, r2start int) {
	//The length of R1 is adjusted to be at least three.
	r1start = romance.VnvSuffix(word, isLowerVowel, 0)
	if r1start < 3 && len(word.RS) >= 3 {
		r1start = 3
	}
	r2start = romance.VnvSuffix(word, isLowerVowel, r1start)
	return
}

// Replace ß with ss
func replaceSharpS(word *snowballword.SnowballWord) {
	word.RS = []rune(strings.ReplaceAll(word.String(), "ß", "ss"))
}

// Capitalize u and y when inbetween vowels.
func capitalizeUY(word *snowballword.SnowballWord) (numCapitalizations int) {
	for i, r := range word.RS {
		if (r == 'y' || r == 'u') && i < len(word.RS) {
			if isLowerVowel(word.RS[i-1]) && isLowerVowel(word.RS[i+1]) {
				if r == 'y' {
					word.RS[i] = 'Y'
				} else {
					word.RS[i] = 'U'
				}
			}
		}
	}

	return
}
