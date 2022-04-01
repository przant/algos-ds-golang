package reverse_string

func reverseStringStack(s string) (reversed string) {
	stack := make([]rune, 0)
	sRune := make([]rune, 0)

	for _, char := range s {
		stack = append(stack, char)
	}

	stSize := len(stack)

	for char := 0; char < stSize; char++ {
		sRune = append(sRune, stack[stSize-(char+1)])
		stack = stack[:stSize-(char+1)]
	}

	return string(sRune)
}

func reverseString(s string) (reversed string) {
	sRune := []rune(s)

	left := 0
	right := len(s) - 1

	for left <= right {
		sRune[left], sRune[right] = sRune[right], sRune[left]
		left++
		right--
	}

	return string(sRune)
}
