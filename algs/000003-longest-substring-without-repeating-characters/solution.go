package longestsubstringwithoutrepeatingcharacters

func Solution(s string) int {
	var (
		left  int
		right int
		res   int
	)
	m := make(map[byte]int)
	for right < len(s) {
		if previous, ok := m[s[right]]; ok {
			left = max(left, previous)
		}
		res = max(res, right-left+1)
		m[s[right]] = right + 1
		right++
	}
	return res
}
