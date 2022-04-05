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
