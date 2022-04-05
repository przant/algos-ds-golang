package valid_palindrome

import "unicode"

func validPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	for start <= end {
		if unicode.ToUpper(rune(s[start])) != unicode.ToUpper(rune(s[end])) {
			return checkSubStr(start+1, end, s) || checkSubStr(start, end-1, s)
		}
		start++
		end--
	}
	return true
}

func checkSubStr(s, e int, subStr string) bool {
	for s <= e {
		if unicode.ToUpper(rune(subStr[s])) != unicode.ToUpper(rune(subStr[e])) {
			return false
		}
		s++
		e--
	}
	return true
}
