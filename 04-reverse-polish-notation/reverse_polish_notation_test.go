package reverse_polish_notation

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{
			input: []string{"20"},
			want:  20,
		},
		{
			input: []string{"2", "1", "+", "3", "*"},
			want:  9,
		},
		{
			input: []string{"4", "13", "5", "/", "+"},
			want:  6,
		},
		{
			input: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:  22,
		},
	}

	for _, test := range tests {
		if got := evalRPN(test.input); got != test.want {
			t.Errorf("with input %q, want %d but got %d\n", test.input, test.want, got)
		}
	}
}
