package longestsubstringwithoutrepeatingcharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc     string
		s        string
		expected int
	}{
		{
			desc:     "Case#1",
			s:        "abcabcbb",
			expected: 3,
		},
		{
			desc:     "Case#2",
			s:        "bbbbb",
			expected: 1,
		},
		{
			desc:     "Case#3",
			s:        "pwwkew",
			expected: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Solution(tC.s)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
