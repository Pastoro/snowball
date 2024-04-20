package german

import "github.com/kljensen/snowball/snowballword"

func step3(w *snowballword.SnowballWord) bool {
	dSuffixes := []string{
		"end", "ung", "ig", "ik", "isch", "lich", "heit", "keit",
	}
	suffix, _ := w.FirstSuffixIfIn(w.R1start, len(w.RS), dSuffixes...)
	if suffix == "end" || suffix == "ung" {
		if w.String()[len(w.RS)-5:len(w.RS)-3] == "ig" && w.String()[len(w.RS)-6] != 'e' && len(w.RS)-5 >= w.R2start {
			w.RemoveLastNRunes(5)
		} else {
			w.RemoveLastNRunes(3)
		}
	} else if suffix == "ig" || suffix == "ik" || suffix == "isch" {
		if w.String()[len(w.RS)-len(suffix)] != 'e' {
			w.RemoveLastNRunes(len(suffix))
		}
	} else if suffix == "lich" || suffix == "heit" {
		if w.String()[len(w.RS)-6:len(w.RS)-4] == "er" || w.String()[len(w.RS)-6:len(w.RS)-4] == "en" {
			w.RemoveLastNRunes(6)
		}
	}
	return true
}
