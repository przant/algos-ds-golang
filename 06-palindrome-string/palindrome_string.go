package palindrome_string

import "unicode"

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start <= end {

		for start < end && !unicode.IsLetter(rune(s[start])) {
			start++
		}

		for start < end && !unicode.IsLetter(rune(s[end])) {
			end--
		}

		if unicode.ToUpper(rune(s[start])) != unicode.ToUpper(rune(s[end])) {
			return false
		}
		start++
		end--
	}
	return true
}
