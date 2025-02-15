package leetcode

import "testing"

func TestLargestOddNumber(t *testing.T) {
	var cases = []struct {
		input  string
		output string
	}{
		{
			input:  "52",
			output: "5",
		},
		{
			input:  "4206",
			output: "",
		},
		{
			input:  "35427",
			output: "35427",
		},
	}
	for _, c := range cases {
		x := largestOddNumber(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}
