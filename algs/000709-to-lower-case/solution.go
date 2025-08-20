package tolowercase

/*
	Given a string s, return the string after replacing every uppercase letter with the same lowercase letter.


	Example 1:

	Input: s = "Hello"
	Output: "hello"


	Example 2:

	Input: s = "here"
	Output: "here"


	Example 3:

	Input: s = "LOVELY"
	Output: "lovely"
*/

// string.ToLower(s)
func Solution(s string) string {
	r := []rune(s)

	for i := range r {
		if isUppercaseLetter(r[i]) {
			r[i] += 32
		}
	}

	return string(r)
}

func isUppercaseLetter(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}
