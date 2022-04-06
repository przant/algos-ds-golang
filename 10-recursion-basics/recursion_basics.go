package recursion_basics

import (
	"fmt"
	"io"
	"math"
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

func numOfDigits(n int) int {
	if n == 0 {
		return 0
	}

	smallAns := numOfDigits(n / 10)

	return smallAns + 1
}

func sumOfDigits(n int) int {
	if n == 0 {
		return 0
	}

	smallAns := sumOfDigits(n / 10)

	return smallAns + n%10
}

func product(n, m int) int {
	if m == 1 {
		return n
	}

	smallAns := product(n, m-1)

	return smallAns + n

}

func countZeroes(n int) int {
	if n == 0 {
		return 0
	}

	count := countZeroes(n / 10)

	if n%10 == 0 {
		return count + 1
	} else {
		return count
	}
}

func geoSum(k int) float64 {
	if k == 0 {
		return 1.0
	}

	smallAns := geoSum(k - 1)

	return smallAns + (1.0 / math.Pow(2, float64(k)))
}
