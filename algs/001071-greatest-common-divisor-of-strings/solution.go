package greatestcommondivisorofstrings

/*
	For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).
	Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.


	Example 1:

	Input: str1 = "ABCABC", str2 = "ABC"
	Output: "ABC"


	Example 2:

	Input: str1 = "ABABAB", str2 = "ABAB"
	Output: "AB"


	Example 3:

	Input: str1 = "LEET", str2 = "CODE"
	Output: ""
*/

func Solution(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	a, b := len(str1), len(str2)
	for b > 0 {
		a, b = b, a%b
	}

	return str1[:a]
}

// func Solution(str1 string, str2 string) string {
// 	maxStr := str1
// 	if len(str1) < len(str2) {
// 		maxStr = str2
// 	}

// 	var (
// 		sb     strings.Builder
// 		result string
// 	)

// 	for _, letter := range maxStr {
// 		sb.WriteRune(letter)

// 		ratio1 := len(str1) / len(sb.String())
// 		ratio2 := len(str2) / len(sb.String())

// 		common := sb.String()

// 		var res1, res2 strings.Builder

// 		for range ratio1 {
// 			res1.WriteString(common)
// 		}

// 		for range ratio2 {
// 			res2.WriteString(common)
// 		}

// 		if res1.String() == str1 && res2.String() == str2 {
// 			result = common
// 		}
// 	}

// 	return result
// }
