package leetcode

import "testing"

func TestRemoveDuplicatesfromSortedArray(t *testing.T) {
	var cases = []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 1, 2},
			output: 2,
		},
		{
			input:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			output: 5,
		},
	}

	for _, c := range cases {
		got := RemoveDuplicatesfromSortedArray(c.input)
		t.Log(got, c.output)
		if got != c.output {
			t.Logf("failed: got %d want %d in: %v\n", got, c.output, c.input)
			t.Fail()
		}
	}
}
