package german

import (
	"strings"

	"github.com/kljensen/snowball/snowballword"
)

func Stem(word string, stemStopWords bool) string {
	word = strings.ToLower(strings.TrimSpace(word))

	if len(word) <= 2 || (!stemStopWords && IsStopWord(word)) {
		return word
	}

	w := snowballword.New(word)
	preprocess(w)
	step1(w)
	step2(w)
	//step3(w)
	return w.String()
}
