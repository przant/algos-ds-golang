package recursion_basics

import (
	"fmt"
	"io"
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	smallAns := factorial(n - 1)

	return n * smallAns
}

func fibonacci(n int) int {
	if n == 2 || n == 1 {
		return 1
	}

	smallAns := fibonacci(n-1) + fibonacci(n-2)

	return smallAns
}

func power(x, n int) int {
	if n == 0 {
		return 1
	}

	smallAns := power(x, n-1)

	return x * smallAns
}

func printNumbersAsc(n int, w io.Writer) {
	if n == 0 {
		return
	}

	printNumbersAsc(n-1, w)

	fmt.Fprintf(w, "%d", n)
}

func printNumbersDesc(n int, w io.Writer) {
	if n == 0 {
		return
	}

	fmt.Fprintf(w, "%d", n)

	printNumbersDesc(n-1, w)

}
