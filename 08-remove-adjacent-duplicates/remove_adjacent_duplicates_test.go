package remove_duplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "aa",
			want:  "",
		},
		{
			input: "abbaca",
			want:  "ca",
		},
		{
			input: "azxxzy",
			want:  "ay",
		},
		{
			input: "aaabccddd",
			want:  "abd",
		},
		{
			input: "Mississippi",
			want:  "M",
		},
	}

	for _, test := range tests {
		if got := removeDuplicates(test.input); got != test.want {
			t.Errorf("with input %q, want %q but got %q", test.input, test.want, got)
		}
	}
}

func TestRemoveDuplicatesPlace(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "aa",
			want:  "",
		},
		{
			input: "abbaca",
			want:  "ca",
		},
		{
			input: "azxxzy",
			want:  "ay",
		},
		{
			input: "aaabccddd",
			want:  "abd",
		},
		{
			input: "Mississippi",
			want:  "M",
		},
	}

	for _, test := range tests {
		if got := removeDuplicatesPlace([]rune(test.input)); got != test.want {
			t.Errorf("with input %q, want %q but got %q", test.input, test.want, got)
		}
	}
}
