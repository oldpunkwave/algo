package mostcommonword

import (
	"strings"
)

/*
	Given a string paragraph and a string array of the banned words banned, return the most frequent word that is not banned.
	It is guaranteed there is at least one word that is not banned, and that the answer is unique.
	The words in paragraph are case-insensitive and the answer should be returned in lowercase.
	Note that words can not contain punctuation symbols.



	Example 1:

	Input: paragraph = "Bob hit a ball, the hit BALL flew far after it was hit.", banned = ["hit"]
	Output: "ball"

	Explanation:
	"hit" occurs 3 times, but it is a banned word.
	"ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph.
	Note that words in the paragraph are not case sensitive,
	that punctuation is ignored (even if adjacent to words, such as "ball,"),
	and that "hit" isn't the answer even though it occurs more because it is banned.


	Example 2:

	Input: paragraph = "a.", banned = []
	Output: "a"
*/

func Solution(paragraph string, banned []string) string {
	const symbols = "!?',;."

	for _, s := range symbols {
		paragraph = strings.ReplaceAll(paragraph, string(s), " ")
	}

	words := make(map[string]int)

	paragraph = strings.ToLower(paragraph)
	for _, v := range strings.Split(paragraph, " ") {
		if v != "" {
			words[v]++
		}
	}

	for _, bannedWord := range banned {
		delete(words, bannedWord)
	}

	var (
		res string
		max int
	)
	for word, counter := range words {
		if counter > max {
			res = word
			max = counter
		}
	}

	return res
}
