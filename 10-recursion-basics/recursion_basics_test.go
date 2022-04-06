package recursion_basics

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{
			input: 0,
			want:  1,
		},
		{
			input: 1,
			want:  1,
		},
		{
			input: 4,
			want:  24,
		},
		{
			input: 5,
			want:  120,
		},
		{
			input: 6,
			want:  720,
		},
	}

	for _, test := range tests {
		if got := factorial(test.input); got != test.want {
			t.Errorf("with input %d, want %d but got %d", test.input, test.want, got)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{
			input: 1,
			want:  1,
		},
		{
			input: 2,
			want:  1,
		},
		{
			input: 3,
			want:  2,
		},
		{
			input: 5,
			want:  5,
		},
		{
			input: 6,
			want:  8,
		},
	}

	for _, test := range tests {
		if got := fibonacci(test.input); got != test.want {
			t.Errorf("with input %d, want %d but got %d", test.input, test.want, got)
		}
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		input [2]int // [2]int{x, n}
		want  int
	}{
		{
			input: [2]int{5, 0},
			want:  1,
		},
		{
			input: [2]int{5, 1},
			want:  5,
		},
		{
			input: [2]int{5, 2},
			want:  25,
		},
		{
			input: [2]int{7, 4},
			want:  2401,
		},
		{
			input: [2]int{10, 4},
			want:  10000,
		},
	}

	for _, test := range tests {
		if got := power(test.input[0], test.input[1]); got != test.want {
			t.Errorf("with input %v, want %d but got %d", test.input, test.want, got)
		}
	}
}
