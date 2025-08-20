package firstbadversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc         string
		input        int
		expected     int
		isBadVersion func(int) bool
	}{
		{
			desc:     "Case#1",
			input:    5,
			expected: 4,
			isBadVersion: func(i int) bool {
				return i >= 4
			},
		},
		{
			desc:     "Case#2",
			input:    1,
			expected: 1,
			isBadVersion: func(i int) bool {
				return i >= 4
			},
		},
		{
			desc:     "Case#3",
			input:    100,
			expected: 6,
			isBadVersion: func(i int) bool {
				return i >= 6
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Solution(tC.input, tC.isBadVersion)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
