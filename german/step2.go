package german

import "github.com/kljensen/snowball/snowballword"

func step2(w *snowballword.SnowballWord) bool {
	suffixes := []string{
		"en", "er", "est", "st",
	}
	suffix, suffixRunes := w.FirstSuffixIfIn(w.R1start, len(w.RS), suffixes...)
	if suffix == "st" {
		rsLen := len(w.RS)
		if rsLen >= 6 {
			//If the valid st-ending is itself preceded by at least 3 letters
			if isValidStEnding(w.RS[len(w.RS)-3]) {
				w.RemoveLastNRunes(len(suffixRunes))
			}
		}
	} else if suffix == "en" || suffix == "er" || suffix == "est" {
		w.RemoveLastNRunes(len(suffixRunes))
	} else {
		return false
	}

	return true
}
