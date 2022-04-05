package excel_column_number

func titleToNumber(s string) (result int) {
	pow := 1
	base := 26
	size := len(s) - 1

	for pos := size; pos >= 0; pos-- {
		result += int(s[pos]-'A'+1) * pow
		pow *= base
	}
	return
}
