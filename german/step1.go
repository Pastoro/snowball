package german

import (
	"github.com/kljensen/snowball/snowballword"
)

func step1(w *snowballword.SnowballWord) bool {

	suffixes := []string{
		"system", "ern", "em", "er", "e", "en", "es", "s",
	}
	suffix, suffixRunes := w.FirstSuffixIn(w.R1start, len(w.RS), suffixes...)
	if suffix == "s" {
		rsLen := len(w.RS)
		if isValidSEnding(w.RS[rsLen-2]) {
			w.RemoveLastNRunes(len(suffixRunes))
		}
	} else if suffix == "system" {

	} else if suffix == "" {
		return false
	} else if suffix == "em" || suffix == "ern" || suffix == "er" || suffix == "e" || suffix == "en" || suffix == "es" {
		w.RemoveLastNRunes(len(suffix))
		//If the suffix is preceded by "niss", remove one s
		if (len(w.RS)) >= 4 && w.String()[len(w.RS)-3:] == "niss" {
			w.RemoveLastNRunes(1)
		}
	}
	return true
}
