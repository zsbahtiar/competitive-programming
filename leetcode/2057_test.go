package leetcode

import "testing"

func TestSmallestIndexWithEqualValue(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{0, 1, 2},
			output: 0,
		},
		{
			input:  []int{4, 3, 2, 1},
			output: 2,
		},
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			output: -1,
		},
	}

	for _, test := range cases {
		got := SmallestIndexWithEqualValue(test.input)
		if got != test.output {
			t.Logf("failed: got %d want %d in: %v", got, test.output, test.input)
			t.Fail()
		}
	}
}
