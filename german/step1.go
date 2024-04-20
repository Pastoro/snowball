package german

import (
	"github.com/kljensen/snowball/snowballword"
)

func step1(w *snowballword.SnowballWord) bool {

	suffixes := []string{
		"ern", "em", "er", "e", "en", "es", "s",
	}
	suffix, suffixRunes := w.FirstSuffixIn(w.R1start, len(w.RS), suffixes...)
	if suffix == "s" {
		rsLen := len(w.RS)
		if rsLen >= 2 {
			if isValidSEnding(w.RS[rsLen-2]) {
				w.RemoveLastNRunes(len(suffixRunes))
			}
		}
	} else {
		w.RemoveLastNRunes(len(suffixRunes))
		if (len(w.RS)) >= 4 && (suffix == "e" || suffix == "en" || suffix == "es") && w.String()[len(w.RS)-3:] == "niss" {
			w.RemoveLastNRunes(1)
		}
	}
	return true
}
