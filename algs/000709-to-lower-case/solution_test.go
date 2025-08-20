package tolowercase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "Case#1",
			input:    "HeLlO",
			expected: "hello",
		},
		{
			desc:     "Case#2",
			input:    "HELLO",
			expected: "hello",
		},
		{
			desc:     "Case#3",
			input:    "hello",
			expected: "hello",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Solution(tC.input)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
