package recursion_basics

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	smallAns := factorial(n - 1)

	return n * smallAns
}
