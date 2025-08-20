package sortarraybyparity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []int
		expected []int
	}{
		{
			desc:     "Case#1",
			input:    []int{3, 1, 2, 4},
			expected: []int{2, 4, 3, 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Solution(tC.input)
			assert.ElementsMatch(t, actual, tC.expected)
		})
	}
}
