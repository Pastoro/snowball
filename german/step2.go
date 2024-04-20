package german

import "github.com/kljensen/snowball/snowballword"

func step2(w *snowballword.SnowballWord) bool {
	suffixes := []string{
		"en", "er", "est", "st",
	}
	suffix, suffixRunes := w.FirstSuffixIfIn(0, w.R2start, suffixes...)
	if suffix == "st" {
		rsLen := len(w.RS)
		if rsLen >= 6 {
			if isValidStEnding(w.RS[len(w.RS)-3]) {
				w.RemoveLastNRunes(2)
			}
		}
	} else if suffix == "en" || suffix == "er" || suffix == "est" {
		w.RemoveLastNRunes(len(suffixRunes))
	}

	return true
}
