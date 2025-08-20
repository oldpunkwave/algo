package robotreturntoorigin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected assert.BoolAssertionFunc
	}{
		{
			desc:     "Case#1",
			input:    "UD",
			expected: assert.True,
		},
		{
			desc:     "Case#1",
			input:    "LL",
			expected: assert.False,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Solution(tC.input)
			tC.expected(t, actual)
		})
	}
}
