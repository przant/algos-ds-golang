package remove_duplicates

import "testing"

func TestRemoveKDuplicates(t *testing.T) {
	tests := []struct {
		kDups    int
		inputStr string
		want     string
	}{
		{
			kDups:    2,
			inputStr: "Mississippi",
			want:     "M",
		},
		{
			kDups:    3,
			inputStr: "ddaabbaaabcccaee",
			want:     "ddee",
		},
		{
			kDups:    4,
			inputStr: "ffddaabbeecccceebbaaddf",
			want:     "fff",
		},
	}

	for _, test := range tests {
		if got := removeKDuplicates(test.kDups, test.inputStr); got != test.want {
			t.Errorf("with input:%d,%q; want %q but got %q",
				test.kDups, test.inputStr, test.want, got,
			)
		}
	}
}
