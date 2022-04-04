package excel_column_number

import "testing"

func TestTitleToNumber(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "A",
			want:  1,
		},
		{
			input: "AA",
			want:  27,
		},
		{
			input: "AAA",
			want:  703,
		},
		{
			input: "ABC",
			want:  731,
		},
	}

	for _, test := range tests {
		if got := titleToNumber(test.input); got != test.want {
			t.Errorf("with input %q, want %d but got %d", test.input, test.want, got)
		}
	}
}
