package pathcrossing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected assert.BoolAssertionFunc
	}{
		{
			desc:     "Case#1",
			input:    "NES",
			expected: assert.False,
		},
		{
			desc:     "Case#2",
			input:    "NESWW",
			expected: assert.True,
		},
		{
			desc:     "Case#3",
			input:    "NESW",
			expected: assert.True,
		},
		{
			desc:     "Case#4",
			input:    "SN",
			expected: assert.True,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Solution(tC.input)
			tC.expected(t, actual)
		})
	}
}
