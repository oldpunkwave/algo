package mintimetotype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected int
	}{
		{
			desc:     "Case#1",
			input:    "bza",
			expected: 7,
		},
		{
			desc:     "Case#2",
			input:    "zjpc",
			expected: 34,
		},
		{
			desc:     "Case#3",
			input:    "abc",
			expected: 5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Solution(tC.input)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
