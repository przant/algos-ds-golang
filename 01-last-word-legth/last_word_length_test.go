package last_word

import "testing"

func TestLastWordLength(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "",
			want:  0,
		},
		{
			input: "     ",
			want:  0,
		},
		{
			input: "a",
			want:  1,
		},
		{
			input: "Hello World",
			want:  5,
		},
		{
			input: "Hello you!",
			want:  3,
		},
		{
			input: "Hello you!     ",
			want:  3,
		},
	}

	for _, test := range tests {
		if got := lastWordLength(test.input); got != test.want {
			t.Errorf("with input %q, want %d but got %d\n", test.input, test.want, got)
		}
	}
}
