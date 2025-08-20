package findthenumberofgoodpairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}

	testCases := []struct {
		desc     string
		args     args
		expected int
	}{
		{
			desc: "Case#1",
			args: args{
				nums1: []int{1, 3, 4},
				nums2: []int{1, 3, 4},
				k:     1,
			},
			expected: 5,
		},
		{
			desc: "Case#1",
			args: args{
				nums1: []int{1, 2, 4, 12},
				nums2: []int{2, 4},
				k:     3,
			},
			expected: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Solution(tC.args.nums1, tC.args.nums2, tC.args.k)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
