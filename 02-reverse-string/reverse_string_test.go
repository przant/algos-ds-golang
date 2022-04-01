package reverse_string

import "testing"

func TestReversedStringStack(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "",
			want:  "",
		},

		{
			input: "abc",
			want:  "cba",
		},
		{
			input: "abc   def",
			want:  "fed   cba",
		},
	}

	for _, test := range tests {
		if got := reverseStringStack(test.input); got != test.want {
			t.Errorf("with input %q, want %q but got %q\n", test.input, test.want, got)
		}
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "",
			want:  "",
		},

		{
			input: "abc",
			want:  "cba",
		},
		{
			input: "abc   def",
			want:  "fed   cba",
		},
	}

	for _, test := range tests {
		if got := reverseString(test.input); got != test.want {
			t.Errorf("with input %q, want %q but got %q\n", test.input, test.want, got)
		}
	}
}
