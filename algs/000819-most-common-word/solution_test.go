package mostcommonword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	type args struct {
		paragraph string
		banned    []string
	}

	testCases := []struct {
		desc     string
		args     args
		expected string
	}{
		{
			desc: "Case#1",
			args: args{
				paragraph: "Bob hit a ball, the hit BALL flew far after it was hit.",
				banned:    []string{"hit"},
			},
			expected: "ball",
		},
		{
			desc: "Case#2",
			args: args{
				paragraph: "a.",
				banned:    []string{},
			},
			expected: "a",
		},
		{
			desc: "Case#3",
			args: args{
				paragraph: "a, a, a, a, b,b,b,c, c",
				banned:    []string{"a"},
			},
			expected: "b",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Solution(tC.args.paragraph, tC.args.banned)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
