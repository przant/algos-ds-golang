package palindrome_string

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "abc",
			want:  false,
		},
		{
			input: "race a car",
			want:  false,
		},
		{
			input: "Oso",
			want:  true,
		},
		{
			input: "A man, a plan, a canal: Panama",
			want:  true,
		},
	}

	for _, test := range tests {
		if got := isPalindrome(test.input); got != test.want {
			t.Errorf("with input %q, want %t but got %t", test.input, test.want, got)
		}
	}
}
