package longest_common_prefix

import "sort"

func longestCommonPrefix(strs []string) (lcp string) {
	sTmpRune := make([]rune, 0)

	if len(strs) == 0 {
		return
	}

	sort.SliceStable(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	arraySize := len(strs)
	minStrSize := len(strs[0])

	for col := 0; col < minStrSize; col++ {
		exitLoop := false
		for row := 0; row < arraySize-1; row++ {
			if strs[row][col] != strs[row+1][col] {
				exitLoop = true
				break
			}
		}

		if exitLoop {
			break
		} else {
			sTmpRune = append(sTmpRune, rune(strs[0][col]))
		}

	}

	lcp = string(sTmpRune)

	return
}
