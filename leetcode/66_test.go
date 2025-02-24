package leetcode

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	var cases = []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{1, 2, 3},
			output: []int{1, 2, 4},
		},
		{
			input:  []int{4, 3, 2, 1},
			output: []int{4, 3, 2, 2},
		},
		{
			input:  []int{9},
			output: []int{1, 0},
		},
	}

	for _, c := range cases {
		got := PlusOne(c.input)
		if !reflect.DeepEqual(got, c.output) {
			t.Logf("failed: got %v want %v in: %v\n", got, c.output, c.input)
			t.Fail()
		}
	}
}
