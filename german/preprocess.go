package german

import "github.com/kljensen/snowball/snowballword"

func preprocess(word *snowballword.SnowballWord) {
	replaceSharpS(word)
	capitalizeUY(word)

	r1start, r2start := r1r2(word)
	word.R1start = r1start
	word.R2start = r2start
}
