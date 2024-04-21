package german

import (
	"github.com/kljensen/snowball/snowballword"
)

func step3(w *snowballword.SnowballWord) bool {
	dSuffixes := []string{
		"isch", "lich", "heit", "keit", "end", "ung", "ig", "ik",
	}

	suffix, suffixRunes := w.FirstSuffixIfIn(w.R2start, len(w.RS), dSuffixes...)
	if suffix == "end" || suffix == "ung" {
		igSuffix, _ := w.FirstSuffixIfIn(w.R2start, len(w.RS)-3, "ig")
		if igSuffix != "" && w.String()[len(w.RS)-5] != 'e' {
			w.RemoveLastNRunes(len(suffix) + len(igSuffix))
		} else {
			w.RemoveLastNRunes(len(suffix))
		}
	} else if suffix == "ig" || suffix == "ik" || suffix == "isch" {
		if w.String()[len(w.RS)-len(suffix)] != 'e' {
			w.RemoveLastNRunes(len(suffix))
		}
	} else if suffix == "lich" || suffix == "heit" {
		if w.String()[len(w.RS)-6:len(w.RS)-4] == "er" || w.String()[len(w.RS)-6:len(w.RS)-4] == "en" && len(w.RS)-6 >= w.R1start {
			w.RemoveLastNRunes(len(suffixRunes) + 2)
		} else {
			w.RemoveLastNRunes(len(suffixRunes))
		}
	} else if suffix == "keit" {
		extraLength := 0
		if w.String()[len(w.RS)-8:len(w.RS)-4] == "lich" {
			extraLength = 4
		} else if w.String()[len(w.String())-6:len(w.String())-4] == "ig" {
			extraLength = 2
		}
		if extraLength != 0 {
			w.RemoveLastNRunes(len(suffix) + extraLength)
		} else {
			w.RemoveLastNRunes(len(suffix))
		}
	}
	return true
}
