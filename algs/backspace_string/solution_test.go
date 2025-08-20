package backspacestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		inputS string
		inputT string
		check  assert.BoolAssertionFunc
	}{
		{
			desc:   "Case#1",
			inputS: "ab##########c",
			inputT: "ad##c",
			check:  assert.True,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Solution(tC.inputS, tC.inputT)
			tC.check(t, actual)
		})
	}

}
