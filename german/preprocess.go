package german

import "github.com/kljensen/snowball/snowballword"

func preprocess(word *snowballword.SnowballWord) {
	capitalizeUY(word)
	mapCharacters(word)

	r1start, r2start := r1r2(word)
	word.R1start = r1start
	word.R2start = r2start
}
