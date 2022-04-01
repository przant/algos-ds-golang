package longest_common_prefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{
			input: []string{""},
			want:  "",
		},
		{
			input: []string{"a"},
			want:  "a",
		},
		{
			input: []string{"abc"},
			want:  "abc",
		},
		{
			input: []string{"abcdef", "", "abc"},
			want:  "",
		},
		{
			input: []string{"abcdefg", "abcde", "abc", "abccddee"},
			want:  "abc",
		},
	}

	for _, test := range tests {
		if got := longestCommonPrefix(test.input); got != test.want {
			t.Errorf("with input %q, want %q but got %q\n", test.input, test.want, got)
		}
	}
}
