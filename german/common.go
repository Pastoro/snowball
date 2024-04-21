package german

import (
	"strings"

	"github.com/kljensen/snowball/romance"
	"github.com/kljensen/snowball/snowballword"
)

func IsStopWord(word string) bool {
	switch word {
	case "aber", "alle", "allem", "allen", "aller", "alles", "als", "also", "am", "an", "ander", "andere", "anderem", "anderen", "anderer", "anderes", "anderm", "andern", "anderr", "anders", "auch", "auf", "aus",
		"bei", "bin", "bis", "bist",
		"da", "damit", "dann", "der", "den", "des", "dem", "die", "das", "daß", "derselbe", "derselben", "denselben", "desselben", "demselben", "dieselbe", "dieselben", "dasselbe", "dazu", "dein", "deine", "deinem", "deinen", "deiner", "deines", "denn", "derer", "dessen", "dich", "dir", "du", "dies", "diese", "diesem", "diesen", "dieser", "dieses", "doch", "dort", "durch",
		"ein", "eine", "einem", "einen", "einer", "eines", "einig", "einige", "einigem", "einigen", "einiger", "einiges", "einmal", "er", "ihn", "ihm", "es", "etwas", "euer", "eure", "eurem", "euren", "eurer", "eures",
		"für", "gegen", "gewesen", "hab", "habe", "haben", "hat", "hatte", "hatten", "hier", "hin", "hinter", "ich", "mich", "mir", "ihr", "ihre", "ihrem", "ihren", "ihrer", "ihres", "euch", "im", "in", "indem", "ins", "ist",
		"jede", "jedem", "jeden", "jeder", "jedes", "jene", "jenem", "jenen", "jener", "jenes", "jetzt", "kann", "kein", "keine", "keinem", "keinen", "keiner", "keines", "können", "könnte",
		"machen", "man", "manche", "manchem", "manchen", "mancher", "manches", "mein", "meine", "meinem", "meinen", "meiner", "meines", "mit", "muss", "musste", "nach", "nicht", "nichts", "noch", "nun", "nur", "ob", "oder", "ohne",
		"sehr", "sein", "seine", "seinem", "seinen", "seiner", "seines", "selbst", "sich", "sie", "ihnen", "sind", "so", "solche", "solchem", "solchen", "solcher", "solches", "soll", "sollte", "sondern", "sonst", "über", "um", "und", "uns", "unse", "unsem", "unsen", "unser", "unses", "unter", "viel", "vom", "von", "vor", "während", "war", "waren", "warst", "was", "weg", "weil", "weiter", "welche", "welchem", "welchen", "welcher", "welches", "wenn", "werde", "werden", "wie", "wieder", "will", "wir", "wird", "wirst", "wo", "wollen", "wollte", "würde", "würden", "zu", "zum", "zur", "zwar", "zwischen":
		return true
	}
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

func mapCharacters(word *snowballword.SnowballWord) {
	r := strings.NewReplacer(
		"ß", "ss",
		"ae", "ä",
		"oe", "ö",
		"que", "que",
		"ue", "ü",
	)
	word.RS = []rune(r.Replace(word.String()))
}
func removeUmlautAccent(word *snowballword.SnowballWord) {
	r := strings.NewReplacer(
		"ä", "a",
		"ö", "o",
		"ü", "u")
	word.RS = []rune(r.Replace(word.String()))
}

// Capitalize u and y when inbetween vowels.
func capitalizeUY(word *snowballword.SnowballWord) (numCapitalizations int) {
	for i, r := range word.RS {
		if (r == 'y' || r == 'u') && i < len(word.RS)-1 && i > 0 {
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
