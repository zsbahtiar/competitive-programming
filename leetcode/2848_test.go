package leetcode

import "testing"

func TestPointsThatIntersectWithCars(t *testing.T) {
	var cases = []struct {
		input  [][]int
		output int
	}{
		{
			input:  [][]int{{3, 6}, {1, 5}, {4, 7}},
			output: 7,
		},
		{
			input:  [][]int{{1, 3}, {5, 8}},
			output: 7,
		},
	}

	for _, c := range cases {
		got := PointsThatIntersectWithCars(c.input)
		if got != c.output {
			t.Logf("failed: got %d want %d in: %v\n", got, c.output, c.input)
			t.Fail()
		}
	}
}
