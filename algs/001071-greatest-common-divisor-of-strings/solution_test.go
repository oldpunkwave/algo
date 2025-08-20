package greatestcommondivisorofstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}

	testCases := []struct {
		desc     string
		args     args
		expected string
	}{
		{
			desc: "Case#1",
			args: args{
				str1: "ABCABC",
				str2: "ABC",
			},
			expected: "ABC",
		},
		{
			desc: "Case#2",
			args: args{
				str1: "ABABAB",
				str2: "ABAB",
			},
			expected: "AB",
		},
		{
			desc: "Case#3",
			args: args{
				str1: "LEET",
				str2: "CODE",
			},
			expected: "",
		},
		{
			desc: "Case#4",
			args: args{
				str1: "ABABABAB",
				str2: "ABAB",
			},
			expected: "ABAB",
		},
		{
			desc: "Case#5",
			args: args{
				str1: "ABAB",
				str2: "ABABABAB",
			},
			expected: "ABAB",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Solution(tC.args.str1, tC.args.str2)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
