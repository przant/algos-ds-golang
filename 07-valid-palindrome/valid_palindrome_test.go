package valid_palindrome

import "testing"

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "abc",
			want:  false,
		},
		{
			input: "aaba",
			want:  true,
		},
		{
			input: "abca",
			want:  true,
		},
		{
			input: "acba",
			want:  true,
		},
	}

	for _, test := range tests {
		if got := validPalindrome(test.input); got != test.want {
			t.Errorf("with input %q, want %t, but got %t", test.input, test.want, got)
		}
	}
}
