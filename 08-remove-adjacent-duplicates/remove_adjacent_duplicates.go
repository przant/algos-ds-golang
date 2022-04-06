package remove_duplicates

func removeDuplicates(s string) string {
	stack := make([]rune, 0, 256)

	for _, char := range s {
		switch {
		case len(stack) == 0:
			stack = append(stack, char)
		case char != stack[len(stack)-1]:
			stack = append(stack, char)
		case char == stack[len(stack)-1]:
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

func removeDuplicatesPlace(s []rune) string {
	ptr := -1
	size := len(s)

	for pos := 0; pos < size; pos++ {
		if ptr == -1 || s[ptr] != s[pos] {
			ptr++
			s[ptr] = s[pos]
		} else {
			ptr--
		}
	}

	if ptr == -1 {
		return ""
	} else {
		return string(s[:ptr+1])
	}

}
