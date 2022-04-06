package remove_duplicates

type Pair struct {
	Char  rune
	Count int
}

type Stack []Pair

func removeKDuplicates(k int, s string) string {
	stack := make(Stack, 0)

	for _, char := range s {
		topElement := Pair{}

		if len(stack) > 0 {
			topElement = stack[len(stack)-1]
		}

		switch {
		case len(stack) == 0:
			stack = append(stack, Pair{Char: char, Count: 1})
		case char != topElement.Char:
			stack = append(stack, Pair{Char: char, Count: 1})
		case char == topElement.Char:
			if topElement.Count+1 == k {
				stack = stack[:len(stack)-1]
			} else {
				stack[len(stack)-1].Count++
			}
		}
	}

	result := make([]rune, 0, len(s))

	for _, pair := range stack {
		for times := 1; times <= pair.Count; times++ {
			result = append(result, pair.Char)
		}
	}

	return string(result)
}
