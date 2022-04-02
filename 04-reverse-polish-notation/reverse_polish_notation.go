package reverse_polish_notation

import (
	"strconv"
	"strings"
)

func evalRPN(tokens []string) int {

	if len(tokens) == 1 {
		result, err := strconv.ParseInt(tokens[0], 10, 64)
		if err != nil {
			panic(err)
		}
		return int(result)
	}

	stack := make([]int, 0)
	digits := "0123456789"

	for _, token := range tokens {
		if strings.ContainsAny(token, digits) {
			result, err := strconv.ParseInt(token, 10, 64)
			if err != nil {
				panic(err)
			}

			stack = append(stack, int(result))
		} else {
			size := len(stack)
			rVal, lVal := stack[size-1], stack[size-2]

			stack = stack[:size-2]

			switch {
			case token == "+":
				stack = append(stack, lVal+rVal)
			case token == "-":
				stack = append(stack, lVal-rVal)
			case token == "*":
				stack = append(stack, lVal*rVal)
			case token == "/":
				stack = append(stack, lVal/rVal)
			}
		}
	}

	return stack[0]
}
