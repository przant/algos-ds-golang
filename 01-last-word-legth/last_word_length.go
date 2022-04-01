package last_word

import "unicode"

func lastWordLength(s string) (count int) {
	pos := len(s) - 1

	for pos >= 0 && !unicode.IsLetter(rune(s[pos])) {
		pos--
	}

	for pos >= 0 && unicode.IsLetter(rune(s[pos])) {
		pos--
		count++
	}
	return
}
