package backspacestring

func Solution(S string, T string) bool {
	sRunes := []rune(S)
	tRunes := []rune(T)
	k := processString(sRunes)
	p := processString(tRunes)

	if k != p {
		return false
	}

	for i := range k {
		if sRunes[i] != tRunes[i] {
			return false
		}
	}

	return true
}

func processString(chars []rune) int {
	k := 0
	for _, c := range chars {
		if c != '#' {
			chars[k] = c
			k++
		} else if k > 0 {
			k--
		}
	}
	return k
}
